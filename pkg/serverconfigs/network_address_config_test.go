package serverconfigs

import "testing"

func TestNetworkAddressConfig_FullAddresses(t *testing.T) {
	{
		addr := &NetworkAddressConfig{
			Protocol:  "http",
			Host:      "127.0.0.1",
			PortRange: "8080",
		}
		err := addr.Init()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(addr.FullAddresses())
	}

	{
		addr := &NetworkAddressConfig{
			Protocol:  "http",
			Host:      "127.0.0.1",
			PortRange: "8080:8090",
		}
		err := addr.Init()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(addr.FullAddresses())
	}

	{
		addr := &NetworkAddressConfig{
			Protocol:  "http",
			Host:      "127.0.0.1",
			PortRange: "8080-8090",
		}
		err := addr.Init()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(addr.FullAddresses())
	}

	{
		addr := &NetworkAddressConfig{
			Protocol:  "http",
			Host:      "127.0.0.1",
			PortRange: "8080-8070",
		}
		err := addr.Init()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(addr.FullAddresses())
	}


	{
		addr := &NetworkAddressConfig{
			Protocol:  "http",
			Host:      "::1",
			PortRange: "8080-8070",
		}
		err := addr.Init()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(addr.FullAddresses())
	}
}

func TestNetworkAddressConfig_PickAddress(t *testing.T) {
	{
		addr := &NetworkAddressConfig{
			Host:      "127.0.0.1",
			PortRange: "1234",
		}
		err := addr.Init()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(addr.PickAddress())
	}

	{
		addr := &NetworkAddressConfig{
			Host:      "127.0.0.1",
			PortRange: "8000-9000",
		}
		err := addr.Init()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(addr.PickAddress())
	}
	{
		addr := &NetworkAddressConfig{
			Host:      "127.0.0.1",
			PortRange: "8000-8001",
		}
		err := addr.Init()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(addr.PickAddress())
	}
	{
		addr := &NetworkAddressConfig{
			Host:      "127.0.0.1",
			PortRange: "9000-8000",
		}
		err := addr.Init()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(addr.PickAddress())
	}
}
