package commands

import (
	"fmt"
	"net"
	"time"

	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"go.uber.org/zap"

	"github.com/doovemax/assh/pkg/config"
)

func cmdPing(c *cli.Context) error {
	if len(c.Args()) < 1 {
		return errors.New("assh: \"ping\" requires exactly 1 argument. See 'assh ping --help'")
	}

	conf, err := config.Open(c.GlobalString("config"))
	if err != nil {
		return errors.Wrap(err, "failed to open configuration file")
	}
	if err = conf.LoadKnownHosts(); err != nil {
		return errors.Wrap(err, "failed to load known-hosts")
	}
	target := c.Args()[0]
	host, err := computeHost(target, c.Int("port"), conf)
	if err != nil {
		return errors.Wrapf(err, "failed to get host %q", target)
	}

	if len(host.Gateways) > 0 {
		return errors.New("assh \"ping\" is not working with gateways (yet)")
	}
	if host.ProxyCommand != "" {
		return errors.New("assh \"ping\" is not working with custom ProxyCommand (yet)")
	}

	portName := "ssh"
	if host.Port != "22" {
		// fixme: resolve port name
		portName = "unknown"
	}
	proto := "tcp"
	fmt.Printf("PING %s (%s) PORT %s (%s) PROTO %s\n", target, host.HostName, host.Port, portName, proto)
	dest := fmt.Sprintf("%s:%s", host.HostName, host.Port)
	count := c.Uint("count")
	transmittedPackets := 0
	receivedPackets := 0
	minRoundtrip := time.Duration(0)
	maxRoundtrip := time.Duration(0)
	totalRoundtrip := time.Duration(0)
	for seq := uint(0); count == 0 || seq < count; seq++ {
		if seq > 0 {
			time.Sleep(time.Duration(c.Float64("wait")) * time.Second)
		}
		start := time.Now()
		conn, err := net.DialTimeout(proto, dest, time.Second*time.Duration(c.Float64("waittime")))
		transmittedPackets++
		duration := time.Since(start)
		totalRoundtrip += duration
		if minRoundtrip == 0 || minRoundtrip > duration {
			minRoundtrip = duration
		}
		if maxRoundtrip < duration {
			maxRoundtrip = duration
		}
		if err == nil {
			defer func() {
				if err2 := conn.Close(); err2 != nil {
					logger().Error("Failed to close connection", zap.Error(err2))
				}
			}()
		}
		if err == nil {
			receivedPackets++
			fmt.Printf("Connected to %s: seq=%d time=%v protocol=%s port=%s\n", host.HostName, seq, duration, proto, host.Port)
			if c.Bool("o") {
				goto stats
			}
		} else {
			// FIXME: switch on error type
			fmt.Printf("Request timeout for seq %d (%v)\n", seq, err)
		}
	}

	// FIXME: catch Ctrl+C

stats:
	fmt.Printf("\n--- %s assh ping statistics ---\n", target)
	packetLossRatio := float64(transmittedPackets-receivedPackets) / float64(transmittedPackets) * 100
	fmt.Printf("%d packets transmitted, %d packets received, %.2f%% packet loss\n", transmittedPackets, receivedPackets, packetLossRatio)
	avgRoundtrip := totalRoundtrip / time.Duration(transmittedPackets)
	fmt.Printf("round-trip min/avg/max = %v/%v/%v\n", minRoundtrip, avgRoundtrip, maxRoundtrip)
	return nil
}
