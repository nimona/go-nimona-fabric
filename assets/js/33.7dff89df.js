(window.webpackJsonp=window.webpackJsonp||[]).push([[33],{400:function(e,t,r){"use strict";r.r(t);var a=r(45),o=Object(a.a)({},(function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[r("h1",{attrs:{id:"hyperspace-v2"}},[r("a",{staticClass:"header-anchor",attrs:{href:"#hyperspace-v2"}},[e._v("#")]),e._v(" Hyperspace v2")]),e._v(" "),r("h2",{attrs:{id:"problem-statement-current-approaches-and-issues"}},[r("a",{staticClass:"header-anchor",attrs:{href:"#problem-statement-current-approaches-and-issues"}},[e._v("#")]),e._v(" Problem statement: Current approaches and issues")]),e._v(" "),r("p",[e._v("So far we've tried two main approaches to peer and content discovery.")]),e._v(" "),r("h3",{attrs:{id:"v0-dht"}},[r("a",{staticClass:"header-anchor",attrs:{href:"#v0-dht"}},[e._v("#")]),e._v(" v0. DHT")]),e._v(" "),r("p",[e._v("The first POC (v0) was a DHT based on Kademlia which is a tried and true\nsolution for many peer-to-peer networks. The main issues with this approach\nwas the high cost (in terms of network requests) of keeping the network up\nto date with what each peer has to offer. In addition to that DHTs make it\nhard to hide what you are offering or what you are looking for.\nThere are a number of attempts at privacy preserving DHTs but they are not\nfun.")]),e._v(" "),r("h2",{attrs:{id:"v1-hyperspace"}},[r("a",{staticClass:"header-anchor",attrs:{href:"#v1-hyperspace"}},[e._v("#")]),e._v(" v1. Hyperspace")]),e._v(" "),r("p",[e._v("The second attempt was a proof-of-concept discovery protocol based around\nusing multi-dimensional vectors created from content hash ngrams.\nEach peer creates an empty vector, splits the hashes of the content they\nwish to advertise in 3 char chunks, hashes them with murmur3 and adds them.\nThey then publish their peer information including this vector to the \"closest\"\npeers. In order to calculate proximity the dot product of each known peer's\nvector is calculated and the outcome the result the closest the peers are.\nFor finding peers, the same process is used in order to create a query vector\nwhich is then sent to the closest peers.\nThis works on a small scale but the network will become unreliable as peers'\nvectors get bigger and the proximity metric starts to become ineffective.")]),e._v(" "),r("p",[e._v("The two main benefits of this approach is that its queries don't expose what\nthe peer is looking for and the fact that peers don't have to spam the network\nwith each new content they want to advertise.\nSince we are using multidimensional vectors it should be relatively easy to\nstart supporting other forms of content discovery that is not based on the\ncontent's hash.")]),e._v(" "),r("p",[r("img",{attrs:{src:"np005-hyperspace-v2-sparse-vector.png",alt:"sparse vector example"}})]),e._v(" "),r("h3",{attrs:{id:"issues"}},[r("a",{staticClass:"header-anchor",attrs:{href:"#issues"}},[e._v("#")]),e._v(" Issues")]),e._v(" "),r("p",[e._v("There is another more inherent problem with both these approaches, and that\nis the fact that they are both pretty demanding in terms of cpu, memory,\nand bandwidth. Peers participating in the network have to constantly be\nsending queries to keep a list of healthy peers, as well as continuously\nadvertise themselves and the content they can provide so they aren't\nforgotten by the network.")]),e._v(" "),r("p",[e._v("This is especially hard for applications on mobile phones or other low\npower/bandwidth devices as actively participating in the network costs both\nbattery and bandwidth.")]),e._v(" "),r("h2",{attrs:{id:"proposal-dense-vector-based-queries-on-a-super-peer-subset"}},[r("a",{staticClass:"header-anchor",attrs:{href:"#proposal-dense-vector-based-queries-on-a-super-peer-subset"}},[e._v("#")]),e._v(" Proposal: Dense vector based queries on a super-peer subset")]),e._v(" "),r("p",[e._v("This proposal aims to continue the work of the current discovery protocol,\nbut instead of trying to find a way to distribute the vectors and queries\nacross all peers on the network, delegate this responsibility to a sub set\nof the network's peers.")]),e._v(" "),r("p",[e._v("The main goals of this update are the following:")]),e._v(" "),r("ul",[r("li",[e._v("Reduce computation and communication cost of the largest percentage of peers")]),e._v(" "),r("li",[e._v("Reduce time required to discover peers providing content")])]),e._v(" "),r("h3",{attrs:{id:"separation-of-concerns"}},[r("a",{staticClass:"header-anchor",attrs:{href:"#separation-of-concerns"}},[e._v("#")]),e._v(" Separation of concerns")]),e._v(" "),r("p",[e._v('Instead of treating all peers as equals, we will be separating them into\ntwo distinct groups based on their capabilities; "hyperspace providers"\nand "hyperspace clients".')]),e._v(" "),r("p",[e._v("Clients will only need to keep a small list of healthy providers and they\nshould be able to trust these won't go away, or at least won't go away often.\nThey also only need to advertise their information and the content they\nprovide to one or more of these providers when either gets updated\n[or if more than x minutes have passed].\n[If the provider does not have an answer to the query it can optionally\nreturn a number of provider the peer can query.]")]),e._v(" "),r("p",[e._v("Providers are responsible for receiving and storing peer announcements\nfor a minimum amount of time in order to be able to respond to queries.\nPartitioning of the vectors across hyperspace providers will be handled in\ndifferent ways throughout the evolution of the protocol and more\ninformation can be found later in the document.")]),e._v(" "),r("p",[e._v("A simpler discover protocol also means it will be easier for developers to\nimplement and debug.")]),e._v(" "),r("h3",{attrs:{id:"hyperspace-v2-0"}},[r("a",{staticClass:"header-anchor",attrs:{href:"#hyperspace-v2-0"}},[e._v("#")]),e._v(" Hyperspace v2.0")]),e._v(" "),r("p",[e._v("In the first version of the network, and while the number of peers and\nadvertised content is small, the providers will be gossiping everything\nbetween them and thus should all have a complete and common understanding\nof the network.")]),e._v(" "),r("p",[e._v("We will also be moving away from creating the vectors from ngrams but rather\nopt for hashing the content id with murmur3 multiple (3 for now) times with\na an incremental seed (0, 1, 2), and adding all hashes to the vector.\nThis makes the vector into something resembling a bloom filter but whether\nit improves anything will have to be seen.")]),e._v(" "),r("h3",{attrs:{id:"hyperspace-v2-1"}},[r("a",{staticClass:"header-anchor",attrs:{href:"#hyperspace-v2-1"}},[e._v("#")]),e._v(" Hyperspace v2.1")]),e._v(" "),r("p",[e._v("In order to move away from the centralised nature of the previous version,\nthe providers will form a structured overlay network that will be used to\npartition the data. There are a number of options on how to go forward with\nthis and additional research will be required.")]),e._v(" "),r("p",[e._v("[In order to reduce the required queries the clients would need, we could\nmake sure that we always replicate the partitioning index to all providers.]")]),e._v(" "),r("h3",{attrs:{id:"messages"}},[r("a",{staticClass:"header-anchor",attrs:{href:"#messages"}},[e._v("#")]),e._v(" Messages")]),e._v(" "),r("ul",[r("li",[r("strong",[e._v("Query")]),e._v("\nAllows looking for peers matching any or all of the bellow:\n"),r("ul",[r("li",[e._v("Public keys")]),e._v(" "),r("li",[e._v("Certificates")]),e._v(" "),r("li",[e._v("Capabilities")]),e._v(" "),r("li",[e._v("Content vector")])])]),e._v(" "),r("li",[r("strong",[e._v("Announce")]),e._v("\nA client informing a provider about their existence or updated information.")]),e._v(" "),r("li",[r("strong",[e._v("Remove")]),e._v("\nA client informing a provider they will be departing the network and that\ntheir information should no longer be advertised.")]),e._v(" "),r("li",[r("strong",[e._v("Report")]),e._v("\nA client informing a provider of another client being unresponsive or\nbehaving badly. This should be better defined.")])]),e._v(" "),r("h2",{attrs:{id:"concerns"}},[r("a",{staticClass:"header-anchor",attrs:{href:"#concerns"}},[e._v("#")]),e._v(" Concerns")]),e._v(" "),r("ul",[r("li",[e._v("Centralisation of discovery")]),e._v(" "),r("li",[e._v("Discovery providers require way more computational resources than normal peers")])]),e._v(" "),r("h2",{attrs:{id:"literature"}},[r("a",{staticClass:"header-anchor",attrs:{href:"#literature"}},[e._v("#")]),e._v(" Literature")]),e._v(" "),r("ul",[r("li",[r("a",{attrs:{href:"https://pdos.csail.mit.edu/~petar/papers/maymounkov-kademlia-lncs.pdf",target:"_blank",rel:"noopener noreferrer"}},[e._v("Kademlia: A peer-to-peer Information System Based on the COR Metric"),r("OutboundLink")],1)]),e._v(" "),r("li",[r("a",{attrs:{href:"https://core.ac.uk/download/pdf/35146789.pdf",target:"_blank",rel:"noopener noreferrer"}},[e._v("Bloofi: Multidimensional Bloom Filters"),r("OutboundLink")],1)]),e._v(" "),r("li",[r("a",{attrs:{href:"http://www.cse.psu.edu/~wul2/Publications/wlee%20icdcs05b.pdf",target:"_blank",rel:"noopener noreferrer"}},[e._v("Supporting Complex Multi-dimensional Queries in P2P Systems"),r("OutboundLink")],1)]),e._v(" "),r("li",[r("a",{attrs:{href:"http://ilpubs.stanford.edu:8090/594/1/2003-33.pdf",target:"_blank",rel:"noopener noreferrer"}},[e._v("Designing a Super-Peer Network"),r("OutboundLink")],1)])])])}),[],!1,null,null,null);t.default=o.exports}}]);