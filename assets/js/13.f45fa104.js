(window.webpackJsonp=window.webpackJsonp||[]).push([[13],{377:function(a,t,s){"use strict";s.r(t);var e=s(44),n=Object(e.a)({},(function(){var a=this,t=a.$createElement,s=a._self._c||t;return s("ContentSlotsDistributor",{attrs:{"slot-key":a.$parent.slotKey}},[s("h1",{attrs:{id:"sonar-testing-tool"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#sonar-testing-tool"}},[a._v("#")]),a._v(" Sonar (Testing tool)")]),a._v(" "),s("p",[a._v('Sonar is a testing tool used as part of our first-pass end to end tests.\nIt allows creating peers that will attempt to "ping" other specified\npeers and wait for them to ping back.')]),a._v(" "),s("p",[a._v("Once sonar has managed to ping all its peers, it then waits to receive pings\nfrom all other defined peers, and exits.")]),a._v(" "),s("h2",{attrs:{id:"env-vars"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#env-vars"}},[a._v("#")]),a._v(" Env vars")]),a._v(" "),s("ul",[s("li",[s("code",[a._v("NIMONA_PEER_PRIVATE_KEY")]),a._v(" - Private key for peer.")]),a._v(" "),s("li",[s("code",[a._v("NIMONA_PEER_BIND_ADDRESS")]),a._v(" - Address (in the "),s("code",[a._v("ip:port")]),a._v(" format) to bind sonar to.")]),a._v(" "),s("li",[s("code",[a._v("NIMONA_PEER_BOOTSTRAPS")]),a._v(" - Bootstrap peers to use (in the "),s("code",[a._v("publicKey@tcps:ip:port")]),a._v("\nshorthand format).")]),a._v(" "),s("li",[s("code",[a._v("NIMONA_SONAR_PING_PEERS")]),a._v(" - Public keys of the peers to lookup and ping.")])]),a._v(" "),s("h2",{attrs:{id:"example"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#example"}},[a._v("#")]),a._v(" Example")]),a._v(" "),s("p",[a._v("Create three peers, one that will be a bootstrap peer.")]),a._v(" "),s("div",{staticClass:"language-txt extra-class"},[s("pre",{pre:!0,attrs:{class:"language-text"}},[s("code",[a._v("Peer 1 (bootstrap):\n  * port: 17000\n  * private: bagacmacaslgzj7n6d4fwxs3by7nuqog45gwm5khjgtrm2eqpntui544swrlxc4cxrivfd2utkafyfkealenybtoxgsbqi5ow3wjfnj3iiyrmcuq\n  * public: bahwqcabaofyfpcrkkhvjgualqkuiawi3qdg5onedar25nxmsk2twqrrcyfja\n\nPeer 2:\n  * port: 17001\n  * private: bagacmacatzmnsq5kjjs2xmmhu6pnruu54uvmonpmjkpjd32sptnszpsejm6hyhbnffcrryogndoa2yhe2g4xcr7stib5w6yggu5yepqvz7mnzwi\n  * public: bahwqcabapqoc2kkfddq4m2g4bvqojunzofd7fgqd3n5qmnj3qi7blt6y3tmq\n\nPeer 3:\n  * port: 17002\n  * private: bagacmacaekai32qmxeol6thr5ml6vpym3mx74zlelifwdhw3smynevgv726mlqvzatc67d5apqkebqdxvggxrfx2ifyfkb53fftqnde5vlctx3i\n  * public: bahwqcabayxblsbgf56h2a7auidahpkmnpclpuqlqkud3wklha2gj3kwfhpwq\n")])])]),s("div",{staticClass:"language-sh extra-class"},[s("pre",{pre:!0,attrs:{class:"language-sh"}},[s("code",[s("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_LOG_LEVEL")]),s("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("error "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),s("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_UPNP_DISABLE")]),s("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("true "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),s("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_PEER_BIND_ADDRESS")]),s("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),s("span",{pre:!0,attrs:{class:"token number"}},[a._v("0.0")]),a._v(".0.0:17000 "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),s("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_PEER_PRIVATE_KEY")]),s("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("bagacmacaslgzj7n6d4fwxs3by7nuqog45gwm5khjgtrm2eqpntui544swrlxc4cxrivfd2utkafyfkealenybtoxgsbqi5ow3wjfnj3iiyrmcuq "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\ngo run ./cmd/bootstrap/main.go\n")])])]),s("div",{staticClass:"language-sh extra-class"},[s("pre",{pre:!0,attrs:{class:"language-sh"}},[s("code",[s("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_LOG_LEVEL")]),s("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("error "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),s("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_UPNP_DISABLE")]),s("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("true "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),s("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_PEER_BIND_ADDRESS")]),s("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),s("span",{pre:!0,attrs:{class:"token number"}},[a._v("0.0")]),a._v(".0.0:17001 "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),s("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_PEER_PRIVATE_KEY")]),s("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("bagacmacatzmnsq5kjjs2xmmhu6pnruu54uvmonpmjkpjd32sptnszpsejm6hyhbnffcrryogndoa2yhe2g4xcr7stib5w6yggu5yepqvz7mnzwi "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),s("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_SONAR_PING_PEERS")]),s("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("bahwqcabayxblsbgf56h2a7auidahpkmnpclpuqlqkud3wklha2gj3kwfhpwq "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),s("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_PEER_BOOTSTRAPS")]),s("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("bahwqcabaofyfpcrkkhvjgualqkuiawi3qdg5onedar25nxmsk2twqrrcyfja@tcps:0.0.0.0:17000 "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\ngo run ./cmd/sonar/main.go\n")])])]),s("div",{staticClass:"language-sh extra-class"},[s("pre",{pre:!0,attrs:{class:"language-sh"}},[s("code",[s("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_LOG_LEVEL")]),s("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("error "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),s("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_UPNP_DISABLE")]),s("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("true "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),s("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_PEER_BIND_ADDRESS")]),s("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),s("span",{pre:!0,attrs:{class:"token number"}},[a._v("0.0")]),a._v(".0.0:17002 "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),s("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_PEER_PRIVATE_KEY")]),s("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("bagacmacaekai32qmxeol6thr5ml6vpym3mx74zlelifwdhw3smynevgv726mlqvzatc67d5apqkebqdxvggxrfx2ifyfkb53fftqnde5vlctx3i "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),s("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_SONAR_PING_PEERS")]),s("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("bahwqcabapqoc2kkfddq4m2g4bvqojunzofd7fgqd3n5qmnj3qi7blt6y3tmq "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),s("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_PEER_BOOTSTRAPS")]),s("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("bahwqcabaofyfpcrkkhvjgualqkuiawi3qdg5onedar25nxmsk2twqrrcyfja@tcps:0.0.0.0:17000 "),s("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\ngo run ./cmd/sonar/main.go\n")])])])])}),[],!1,null,null,null);t.default=n.exports}}]);