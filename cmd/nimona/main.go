package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/profile"
	ishell "gopkg.in/abiosoft/ishell.v2"

	"nimona.io/go/api"
	"nimona.io/go/base58"
	"nimona.io/go/blocks"
	"nimona.io/go/crypto"
	"nimona.io/go/dht"
	"nimona.io/go/net"
	"nimona.io/go/peers"
	"nimona.io/go/storage"
	"nimona.io/go/telemetry"
)

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"

	bootstrapPeerInfos = []string{
		// local
		// andromeda.nimona.io
		"5w1MRKRrbrdhWyb5H3MERKXY2pdzADu12FmxmGFDHvqgucmpVG2R3Z158oseZabUNn14Mq3Mcu5ee14QFLZpeoRbNYw6G5nypbVkpWqbugXRzUcXUL1YBBJhHobr5crJ8RHvj2QmfwEMQav1FdjJgDmXcrZJaqmaG1KdDV9t5R5suhe26LsE2aiihPRcMYfjc6iaPLyKoLpakgDaKrWxx37y9BUJHFcmcyGy3KvPpeWJAhW41Y1xDTwKBhT5ssNwH5NpbE7t27yLDYuGqMqP6gwyL4img9GWuS3eCpvj7ZYBPsGWiKPFpij8netndLGLm9MzE9g9twvWVbpTfhyEnuqXy8FHXQv4PQHP8SjE6vFhw4W96KNQYWgFktTMumcQqhZow6STdnbCZiPCva",
		// borealis.nimona.io
		// "27oJKCRfRfS5ec2Np8nNAb71gqQLJ3UBYqRL8L6tEt6GEmzMopSU5yPoYn8aZUGPaAW3kGtPVjqqqNZYAKE1HQeJwzemhn9U5QAkzF9nYtSvVcCGS7moJLTW4ASq7YJLPCLvR2TrEd5H2gSE16FvrhnMLCGXALwT6ezWHabrtX4upiP852J6ZzTkRVqt1i183L4at2VLDYBgC3G45muBwbuz6HaLgt54Dng5JPhSxj42z8hfwwUUzQuth1K9Y7bbSDi4dCL5PYSovHAJ3AUVbkAoXnXtttewGJ5Fiwcuf6JAPkG3VPPT3trS5bqJRifnDgwXprbw8fX2jmZYyzkF7w8kmYprCJ4rY1fMzt1VAEUr54WAavZqXuPEbYE5YA19RXLtdSseCho7KnNvJ",
		// cassiopeia.nimona.io
		// "NmNZ8LLhvr4yqJmPsoNJELmmy3GFHhvdYaHEJGbQ7UctBMLB9ZhZrMoNBj3ceTEsbBWrFGo5XkpgSCC9SUUPuFBL9qenc2ajwp1R3YoMg82HrWMXPNigfH9BCx6fNYCi2Vhzt6FWB2HsrA7s2o3iXE3HyEHoMRqDJ4NWNcGCDfbLK5kbuYXsLFjnPYDoYBGfXxNRcCvg8ACzDLKZ7xiCWfV83QwsweqKtAYC7hPxMKRQxU2vkG4SaemvjY2jXi6iMyJ8brX73TV3ukvyRuM2fDuYcU1szjF5ro2ZytMLbzpjVAsQ8VXf21AbRLja65CxGnP8jgF97dzuYYbm6ZcGs5vv4xdo467eQM4a3MFXNTpX6L714fqwCnmuLLLbwW9jYg1WCUpuwpW7373m2cW",
		// draco.nimona.io
		// "ki2pmv2fq8AWoqLruGH484CGze7RJkY3TRb7URXZzL2jbXkgoobxrAYj9pkrfgvPa3cniuEDpBVg3CgMDrqmXaSwmro793HBhUppiqhCTHiCFcW1kKD3tj2DuVX8NiyUeEA3H2f4j4kxGc1ff9LNug9kHcAyicUZWpCeYD5GM97XN6dkFZUsqCPzFEC6rtJahxd9zGHxuX49ZQerZVEQ6ppxCNuV4sXDVDbqt2zDuy4PGFSFhoPNmxqNTMqwtfwoKbqDtU4D2LPaagBc54jmyiq177SV84as2BQGapZdxkyXNNkTcTWoB9JigsoqVandvnoJ9s8rqcgv91aD3KqAafEvFRFjMJk61rukGHqpxb4XBRwSX6nfQXzzuRn1833VqGAybN68jjUA",
		// eridanus.nimona.io
		// "27oJKCRfRfS5ec2Np8nNAb71gqQLJ3UBYqRL8L8tMRDoqwxBqrLxKacKBgu2RgUBnipcpayTB11cyNgcFydsmu6f61CNLXMCkEGGVadmManxXKeWXdyvpkLkqwmvd55KSp2kjDCPXdmJv8qeaaZUDMczFXWMbKrDJvBCFRtf3RWjxe4aikoZiysvYRhT8kGtHYZtsGfieNN2acdSJ8ppXWkkr31dQxun6Us3tPwLvaZ4wEAEGd3qqYfErqe2jJ36UqKqBb5pZ2D9Y3ecBPq8jzQ8pVTMkQTSLkMzHKQQkV4cpeJsHDHQrXS5zdaPg4W3y3ShApSMPMhwem83H2QAufEZqdQ5A6ntzxY1chft5YwPW3cqtPnB5VVj91rB27XAxJ18UMtWBuYJBo8ht",
		// fornax.nimona.io
		// "4Kv73sF4NPQczZiNfFgRrn96mbQoLHMAzzv8sc1r4jdA7iJQqhNiP69stvow6Ha4EsxNBMafkjgUHtmrggm1YnAUP1HcbHP4L2G9m5MpFea38jbFTTNyQi3DAetW9PyAejDGVXxsTRjbwf8wvueWLhdEmmwDNqHkjHUM78pdKZ4MmK5NsFWv8tBasMUEAzTCtobei2dfjp5CLMtEq96JDo47h3x2PQa4JFgYkoErXvCymF9gn8t4k2EFV7WyvHku8FMSBNXRA7yWHPByZSK2MFDL7cnxNeoYPRZb1AH9kj1Qyc3TH7MxAuhSS284UG7xAfiXXLYE3F67JSiWtQVUYhGQqc14vu7cioEtoyHkE7siSDxJnhnyE2534k23bJG9eQRytqBQCZ98CN",
		// gemini.nimona.io
		// "4Kv73sF4NPQczZiNfFgRrn96mbQoLHMAzzv8sc3kMkn5FadV7vPmeguPi5ZprjBma5StY4NQ7u6MLF2ogFBKkokX4b6bTv2apBYRmPrN6fWjNCmjwhNYxcQXX3KKfGeeK4asomFRfVhjUMbmtLavvkNH1vEW84oQrCiWsQpHUQBnuwazAY9GgxYr7pLDPZkuUM171SEZze8e6yj7RbP1fpN79ewymREhv5zGTBySNqQz1QZSZzq2MVxYfpW3jsHYhiJhgGDsBc9RfzqtnWYd8mVrPq3U5jFv44rtMztCHi1eyGV46wd6qtjdaAdAJazrcYzTcMVu4EfqnvGf5wgU5xku62oWrxVa2eXFz5D6zm18fgiLznBXdBzVUPXbb231AZPkxsyDBckRki",
		// hydra.nimona.io
		// "ki2pmv2fq8AWoqLruGH484CGze7RJkY3TRb7UTLJkFnAnx1TDX5d4aFKExGwuFJpyMuu9u5WR5SWJAEBbdKEm85PSm2YQ2YpgnWzwT9rjvMKzux4BW4AVWf8B9q1sktgHRUv9SvmWXqqxXxbDWjtNN21YxC2si1rkAw7UVo9BX2KLUwKdvLGPjwFR4cMmKni3TT3h3sKRK4WkiJxmDx5evcriFVJa7pebVUTWsU3SRXK5PBzrNw3MVv6eQUQRh2fXUMiP1hqNHbsspMBhZkMnnkFf83fdbSQu9ZQppXQ3zTP3QNVCAETXb35yTsegZ3ZLbEHdaTdaxWcFbdgqGSf57qLdheCgNRZ9XZYoqcQiM8XGofLRQCdQkVCRJmYznw4jtt15hro8nRL",
		// indus.nimona.io
		// "ki2pmv2fq8AWoqLruGH484CGze7RJkY3TRb7UTmCJ5DiHSQqNwQshJ9vxHt6WS1kzkkJZiJsjWGuLsmACt3qx2RpUE1HJupSqiKvtHkU8Jfpz8of8huHne4oWtz3pBPSWfQzNvvfVqQSaFrrMKbsnShkwybfDV3MdzTwGjYX7x9LZjYmnbaNkN1tp5t2Zq2RKAQuc2X2NUuRssMQt1gRkkDpLTUGWvc4MHNTTYM5qpKt4GpudFop3ZeWExterDurCwMADgAqLY87NxPXayLPTLRN69XB2A8hLJPrzpbyhrFAgnV2tMJ5HNiJxfMaseJhaLidS5aeGkAC6Pq35BdC56sS4z2ALqZycH4qdF6mL9it5UD2R8DJAVRwUKJo2yXYZnwvJJXPVem4",
		// lacerta.nimona.io
		// "FgVPhem3sPuFs5Pcbqo2iiEkTpt3W7A7VHcqbxs4oH5gobYKVW9HNF3TwHksnQDAQJfBftNcGn1EVtBoepvN5GhBtRtnzViGvgimbwCKEhgrXGBNQU3RYGTfET6kPTa5NT4nmezkKSa5tPL3kcVsCYRQuHMdo39VTWGdJ4nLpxz2L3x52uuJXJMvXHz6udrgJvr6fABdW8T2acKNDfuxiLW7WBx1yu4pLE9LhFEiHaAPfasaBWjWArQuxpFQKU5rx8Ut6Q8p1ngjjgnkc4XPUjd2ak5NYpEhS4Vq4E5SbVftZSRyU7csefeDfmUY7pxtAyGWGQuY8qVJYxxvmgBacdJuYZJof6HZ9nf9DGJRcJVB6hoGB9ZA94MpBsKrNRdD715J3UGN6mAsB7C",
		// mensa.nimona.io
		// "ki2pmv2fq8AWoqLruGH484CGze7RJkY3TRb7UVYJwSxt1vHpXFdmK8V46SswswCKtFBJ7KHwzf418MtqgtAKhfgknwAU4UQcw4J8tLBqDuQGkHaJrDRyVdgBPs62bksdWrm71gjpLhDWGJQNBEzZG5YjTZxKQBf9p2pDAQues4j7jx3kpNWTfxfJ3Xh9Y2S2nM1dCVb5AN4mWpLFCy4poC6o4RyWHPgfQ8Ts91aauLnu14v8sfHcjfJvjPteahcEXPYrrA1PfJzG3q3cLq4NdmeK8q8fzfBSpbyfTa3M3Lknfy3rv8PkqvVZBxF4ZJexKTTF1Dqxr93tdCMUpUJt8ULisCEqSKqytmNpjaG1My7J6qAa2n6ohynqtNgUH5M5gVSk93quiFrv",
		// norma.nimona.io
		// "ki2pmv2fq8AWoqLruGH484CGze7RJkY3TRb7UW1LK4CrwSWJyq53cXq4ctKJnSk2DsmbKFHqvEXaRWc4zfQDUXtJS3UvpX8wAqpPbay4ZCw2NiFyoCSNRzz9ZZzh9GZuRaMke8NT6qZGuuay3QY2XqXjHJt7SYJtzeXsWg1ayDrYCRtZ3DHz6SAWL3kNerTmnujpiVowwBdy9FFBcFbSJ1KLVCWkfK58r5gPCCRDVUZnQZUWABCJ1fugJhrYAPQY9crS6jYtZ2YPezPjYapwvQtxZdMCkjrreiEb26Y3ngd32y24FdsawNU6FaM9AJnhnbe5xy5nHyfrkWE2yLdQQvYcAHxLEaoBbfbJchBUiDiRP9dacedbPJU9BoDTwQEr4U2CrSghcr7Q",
		// orion.nimona.io
		// "ki2pmv2fq8AWoqLruGH484CGze7RJkY3TRb7UWTe9BjxnN9s2HFwAFk9SyCQXzApsQpBok6GZg3NaH7TuDi8fMybHuEtFvCRUG72KrDFSDKCfH4PrQxMLns7iJrhdQiG8oNLrUx13MJ7tTMzjAXW2D1RBci5593uCBmV4y7KcRRWXJvGJWNHUuLYChhikLxzhDUd4EKXRJjVwT2TYin84bMCoMsaUusNBnZ5YdznEiFuohg3sRcMwqZQeH53BrounbcKeFDUfd8M6HNW41HwGZWmM47MuqKoLqSPmXm1BaGChu1bejDNNPfWwCYaRB6kRD1zcKHeXzUhefHCbdV4G8vybFnFagwWZF9HVbQaCkMHWH8Jyzn6oQDcpm11hCLYxi62ETpFvSnr",
		// pyxis.nimona.io
		// "ki2pmv2fq8AWoqLruGH484CGze7RJkY3TRb7UWvN6GFsty9vCjjcpqt4jTVbxDvzMLSwY297FFKbmpaZE9X6qrBzjGuV6vD8c7QTGzzEExfT7aUa6yhytiGiti22HdWAWeAvLnnsmMHT66FD2AbjX4dMBJaubNPUUxpxSNSqKSKZPkLK3uHrfKyrSFwP58uVTyby4wKx7ZszVx8Q7y4bvDBCQGwQ8x5nQY5f87hgsGivHufPL3ZgGd1hDPFh6cGLvSKvCRoMgMSYYD5B66HXSH8iNpDQjMmrMcXSgigV9U94s1fQ89seZqDUJ57aRe1fsX4vVJNj1Rx2XqKztAqaRoqr4T6qcKYnbm9j2BMxnxQDP5wHkB5qM2WEovpPeCWXgQkQjfDVKZnn",
		// local.nimona.io
		// "ki2pmv2fq8AWoqLruGH484CGze7RJkY3TRb7UV7KF8YXMYFtEUfJk2TaFzswyJxdGkS4F8swRHAg9xU6MExpex4VMEzrPL6cqZtAdBTZ1pbjueEHpQUGjRPCCgH9jAJ5o8a9dcHEFPRufDfQDxHiEeaK9GVEfigMMusnHC9pz1g11LSymU4gzCUYkNjwY7anCvgfCMrSiGNgMqWjMnFVN4heUCVPjCKHtkAA9hJYFj3dwQrMutJiVkTnCucC8UoCc2YxnpegRbVhzXeJ6tJM1CzdFz4i3WvNEtSLHDEJvKg1GDDCCyNatXX2vqu11aEhj2x4JTKPExpF7WrP1G6ikCmh3YZzvBh3rU1cYvu5sxbzxmkW9tURjW2NtUJrS8rc1Wbd83KFgYNS",
		// stats.nimona.io
		"ki2pmv2fq8AWoqLruGH484CGze7RJkY3TRb7UYFsBgoXDpx8MZD3BaSZG4KD7cbmdghvoyEf9U9d6HPzVTsfinFwVpRhwnyDSw51tVW3vCKnscjdVDBXrkSCwKb97qkSLz31NGgpkgSWYDNNmN48hAePKgk4zVhbzGmBs27s8nWz1eai8TFqqhyoFUMRy5yD5TsqEJy43LqtVsWHhbi6bEQYZFSLc11wbqFkpFt5dmAqQ25dza7jhPXrQx6g1oCD2S5P8qzrd5cxs3iw6pDJMXemFLzeCJRrWrnCasnJGR2rRHjU3yXZrqzSv9RZS7tvbQvjYhH8n2jaaUr6SmSdnSLEUhuer8MzgVy9oUEPqPFv6c4QVEaBbh7tojNjtX2R7B3mmwj6L4ge",
	}
)

func base64ToBytes(s string) []byte {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}

// Hello payload
type Hello struct {
	Body      string            `json:"body"`
	Signature *crypto.Signature `json:"-"`
}

func (h *Hello) GetType() string {
	return "demo.hello"
}

func (h *Hello) GetSignature() *crypto.Signature {
	return h.Signature
}

func (h *Hello) SetSignature(s *crypto.Signature) {
	h.Signature = s
}

func (h *Hello) GetAnnotations() map[string]interface{} {
	// no annotations
	return map[string]interface{}{}
}

func (h *Hello) SetAnnotations(a map[string]interface{}) {
	// no annotations
}

func init() {
	blocks.RegisterContentType(&Hello{}, blocks.Persist())
}

func main() {
	// defer profile.Start(profile.MemProfile).Stop()

	configPath := os.Getenv("NIMONA_PATH")

	if configPath == "" {
		usr, _ := user.Current()
		configPath = path.Join(usr.HomeDir, ".nimona")
	}

	if err := os.MkdirAll(configPath, 0777); err != nil {
		log.Fatal("could not create config dir", err)
	}

	addressBook, err := peers.NewAddressBook(configPath)
	if err != nil {
		log.Fatal("could not load key", err)
	}

	port, _ := strconv.ParseInt(os.Getenv("PORT"), 10, 32)

	// statsBootstrapPeer := &peers.PeerInfo{}
	for _, bootstrapPeer := range bootstrapPeerInfos {
		peerInfoBytes, _ := base58.Decode(bootstrapPeer.key)
		typedPeerInfo, err := blocks.UnpackDecode(peerInfoBytes)
		if err != nil {
			log.Fatal("could not unpack bootstrap node", err.Error())
		}
		peerInfo := typedPeerInfo.(*peers.PeerInfo)
		if err := addressBook.PutPeerInfo(peerInfo); err != nil {
			log.Fatal("could not put bootstrap peer", err)
		}
		// if bootstrapPeer.alias == "stats.nimona.io" {
		// 	statsBootstrapPeer = peerInfo
		// }
		if os.Getenv("RELAY") != "false" {
			addressBook.AddLocalPeerRelay(peerInfo.Thumbprint())
		}
		addressBook.SetAlias(peerInfo.Signature.Key, bootstrapPeer.alias)
	}

	storagePath := path.Join(configPath, "storage")

	dpr := storage.NewDiskStorage(storagePath)
	n, _ := net.NewExchange(addressBook, dpr, fmt.Sprintf("0.0.0.0:%d", port))
	dht, _ := dht.NewDHT(n, addressBook)
	// telemetry.NewTelemetry(n, addressBook.GetLocalPeerKey(),
	// 	statsBootstrapPeer.Signature.Key)

	n.RegisterDiscoverer(dht)

	// n.Listen(context.Background(), fmt.Sprintf("0.0.0.0:%d", port))

	n.Handle("demo.hello", func(payload blocks.Typed) error {
		fmt.Printf("___ Got block %s\n", payload.(*Hello).Body)
		return nil
	})

	httpPort := "26880"
	if nhp := os.Getenv("HTTP_PORT"); nhp != "" {
		httpPort = nhp
	}
	httpAddress := ":" + httpPort
	api := api.New(addressBook, dht, dpr)
	go api.Serve(httpAddress)

	shell := ishell.New()
	shell.Printf("Nimona DHT (%s)\n", version)

	putProvider := &ishell.Cmd{
		Name:    "providers",
		Aliases: []string{"provider"},
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)

			if len(c.Args) < 1 {
				c.Println("Missing providing key")
				return
			}
			key := c.Args[0]
			ctx := context.Background()
			if err := dht.PutProviders(ctx, key); err != nil {
				c.Printf("Could not put key %s\n", key)
				c.Printf("Error: %s\n", err)
			}
		},
		Help: "announce a provided key on the dht",
	}

	getProvider := &ishell.Cmd{
		Name:    "providers",
		Aliases: []string{"provider"},
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)

			if len(c.Args) == 0 {
				c.Println("Missing key")
				return
			}
			c.ProgressBar().Indeterminate(true)
			c.ProgressBar().Start()
			key := c.Args[0]
			ctx := context.Background()
			rs, err := dht.GetProviders(ctx, key)
			c.Println("")
			if err != nil {
				c.Printf("Could not get providers for key %s\n", key)
				c.Printf("Error: %s\n", err)
			}
			providers := []string{}
			for provider := range rs {
				providers = append(providers, provider.Thumbprint())
			}
			c.Println("* " + key)
			for _, peerID := range providers {
				c.Printf("  - %s\n", peerID)
			}
			c.ProgressBar().Stop()
		},
		Help: "get peers providing a value from the dht",
	}

	getBlock := &ishell.Cmd{
		Name:    "blocks",
		Aliases: []string{"block"},
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)

			if len(c.Args) < 1 {
				c.Println("Missing block id")
				return
			}

			ctx, cf := context.WithTimeout(context.Background(), time.Second*10)
			defer cf()

			block, err := n.Get(ctx, c.Args[0])
			if err != nil {
				c.Println(err)
				return
			}

			c.Printf("Received block %#v\n", block)
		},
		Help: "get peers providing a value from the dht",
	}

	listProviders := &ishell.Cmd{
		Name:    "providers",
		Aliases: []string{"provider"},
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)

			ps, _ := dht.GetAllProviders()
			for key, vals := range ps {
				c.Println("* " + key)
				for _, val := range vals {
					c.Printf("  - %s\n", val)
				}
			}
		},
		Help: "list all providers stored in our local dht",
	}

	listPeers := &ishell.Cmd{
		Name:    "peers",
		Aliases: []string{"peer"},
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)

			ps, _ := addressBook.GetAllPeerInfo()
			for _, peer := range ps {
				c.Println("* " + peer.Thumbprint())
				c.Printf("  - addresses:\n")
				for _, address := range peer.Addresses {
					c.Printf("     - %s\n", address)
				}
			}
		},
		Help: "list all peers stored in our local dht",
	}

	listBlocks := &ishell.Cmd{
		Name: "blocks",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)

			blocks, err := dpr.List()
			if err != nil {
				c.Println(err)
				return
			}
			for _, block := range blocks {
				c.Printf("     - %s\n", block)
			}
		},
		Help: "list all blocks in local storage",
	}

	listLocal := &ishell.Cmd{
		Name: "local",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)

			peer := addressBook.GetLocalPeerInfo()
			c.Println("* " + peer.Thumbprint())
			c.Printf("  - addresses:\n")
			for _, address := range peer.Addresses {
				c.Printf("     - %s\n", address)
			}
		},
		Help: "list protocols for local peer",
	}

	send := &ishell.Cmd{
		Name: "send",
		Func: func(c *ishell.Context) {
			if len(c.Args) < 2 {
				c.Println("Missing peer id or block")
				return
			}
			ctx := context.Background()
			msg := strings.Join(c.Args[1:], " ")
			peer, err := addressBook.GetPeerInfo(c.Args[0])
			if err != nil {
				c.Println("Could not get peer")
				return
			}
			signer := addressBook.GetLocalPeerKey()
			if err := n.Send(ctx, &Hello{Body: msg}, peer.Signature.Key, blocks.SignWith(signer)); err != nil {
				c.Println("Could not send block", err)
				return
			}
		},
		Help: "list protocols for local peer",
	}

	block := &ishell.Cmd{
		Name: "block",
		Help: "send blocks to peers",
	}

	get := &ishell.Cmd{
		Name: "get",
		Help: "get resource",
	}

	get.AddCmd(getProvider)
	get.AddCmd(getBlock)

	put := &ishell.Cmd{
		Name: "put",
		Help: "put resource",
	}

	put.AddCmd(putProvider)

	list := &ishell.Cmd{
		Name:    "list",
		Aliases: []string{"l", "ls"},
		Help:    "list cached resources",
	}

	list.AddCmd(listProviders)
	list.AddCmd(listPeers)
	list.AddCmd(listLocal)
	list.AddCmd(listBlocks)

	shell.AddCmd(block)
	shell.AddCmd(get)
	shell.AddCmd(put)
	shell.AddCmd(list)
	shell.AddCmd(send)

	// when started with "exit" as first argument, assume non-interactive execution
	if len(os.Args) > 1 && os.Args[1] == "exit" {
		shell.Process(os.Args[2:]...)
	} else {
		// start shell
		shell.Run()
		// teardown
		shell.Close()
	}
}
