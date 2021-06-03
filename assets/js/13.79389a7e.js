(window.webpackJsonp=window.webpackJsonp||[]).push([[13],{377:function(a,s,t){"use strict";t.r(s);var e=t(48),n=Object(e.a)({},(function(){var a=this,s=a.$createElement,t=a._self._c||s;return t("ContentSlotsDistributor",{attrs:{"slot-key":a.$parent.slotKey}},[t("h1",{attrs:{id:"sonar-testing-tool"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#sonar-testing-tool"}},[a._v("#")]),a._v(" Sonar (Testing tool)")]),a._v(" "),t("p",[a._v('Sonar is a testing tool used as part of our first-pass end to end tests.\nIt allows creating peers that will attempt to "ping" other specified\npeers and wait for them to ping back.')]),a._v(" "),t("p",[a._v("Once sonar has managed to ping all its peers, it then waits to receive pings\nfrom all other defined peers, and exits.")]),a._v(" "),t("h2",{attrs:{id:"env-vars"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#env-vars"}},[a._v("#")]),a._v(" Env vars")]),a._v(" "),t("ul",[t("li",[t("code",[a._v("NIMONA_PEER_PRIVATE_KEY")]),a._v(" - Private key for peer.")]),a._v(" "),t("li",[t("code",[a._v("NIMONA_PEER_BIND_ADDRESS")]),a._v(" - Address (in the "),t("code",[a._v("ip:port")]),a._v(" format) to bind sonar to.")]),a._v(" "),t("li",[t("code",[a._v("NIMONA_PEER_BOOTSTRAPS")]),a._v(" - Bootstrap peers to use (in the "),t("code",[a._v("publicKey@tcps:ip:port")]),a._v("\nshorthand format).")]),a._v(" "),t("li",[t("code",[a._v("NIMONA_SONAR_PING_PEERS")]),a._v(" - Public keys of the peers to lookup and ping.")])]),a._v(" "),t("h2",{attrs:{id:"example"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#example"}},[a._v("#")]),a._v(" Example")]),a._v(" "),t("p",[a._v("Create three peers, one that will be a bootstrap peer.")]),a._v(" "),t("div",{staticClass:"language-txt extra-class"},[t("pre",{pre:!0,attrs:{class:"language-text"}},[t("code",[a._v("Peer 1 (bootstrap):\n  * port: 17000\n  * private: bagacnag4afafnpql33c4aezi6wqduwc45qd3evanoowrz66vjpapoz7zzmakbpnvq2jerpsoygzdincom76dovh7szmo5r5eug5atlk6tsf53hqyz4\n  * public: bahwqdag4aeqllbusjc7e5qnsgq2e4z74g5kp7fsy53d2jin2bgwv5hel3wpbrty\n\nPeer 2:\n  * port: 17001\n  * private: bagacnag4afagf2gpy7rkgni3ytui7kkyfek2mnc25qru4hh35owk4mdkorc632bhv2mglq52q4n3yi6x4nehagak3p2tqjow5ss2zfttatri6xsdvm\n  * public: bahwqdag4aeqcpluymxb3vby3xqr5py2ioamavw7vhas5n3ffvslhgbhcr5pehky\n\nPeer 3:\n  * port: 17002\n  * private: bagacnag4afamw4hk6fk2yqr7fyug4cpf6fs6pzhkyttmwt4m27emmc76sdjcg6o7xfkyfrhw5gs3sx6r5aorj66xnveeo7wpeu456c4vsp3ga4z6ie\n  * public: bahwqdag4aeqn7okvqlcpn2nfxfp5d2a5ct55o3kii57m6jjz34fzle7wmbzt4qi\n")])])]),t("div",{staticClass:"language-sh extra-class"},[t("pre",{pre:!0,attrs:{class:"language-sh"}},[t("code",[t("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_LOG_LEVEL")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("error "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_UPNP_DISABLE")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("true "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_PEER_BIND_ADDRESS")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token number"}},[a._v("0.0")]),a._v(".0.0:17000 "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_PEER_PRIVATE_KEY")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("bagacnag4afafnpql33c4aezi6wqduwc45qd3evanoowrz66vjpapoz7zzmakbpnvq2jerpsoygzdincom76dovh7szmo5r5eug5atlk6tsf53hqyz4 "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\ngo run ./cmd/bootstrap/main.go\n")])])]),t("div",{staticClass:"language-sh extra-class"},[t("pre",{pre:!0,attrs:{class:"language-sh"}},[t("code",[t("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_LOG_LEVEL")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("error "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_UPNP_DISABLE")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("true "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_PEER_BIND_ADDRESS")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token number"}},[a._v("0.0")]),a._v(".0.0:17001 "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_PEER_PRIVATE_KEY")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("bagacnag4afagf2gpy7rkgni3ytui7kkyfek2mnc25qru4hh35owk4mdkorc632bhv2mglq52q4n3yi6x4nehagak3p2tqjow5ss2zfttatri6xsdvm "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_SONAR_PING_PEERS")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("bahwqdag4aeqn7okvqlcpn2nfxfp5d2a5ct55o3kii57m6jjz34fzle7wmbzt4qi "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_PEER_BOOTSTRAPS")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("bahwqdag4aeqllbusjc7e5qnsgq2e4z74g5kp7fsy53d2jin2bgwv5hel3wpbrty@tcps:0.0.0.0:17000 "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\ngo run ./cmd/sonar/main.go\n")])])]),t("div",{staticClass:"language-sh extra-class"},[t("pre",{pre:!0,attrs:{class:"language-sh"}},[t("code",[t("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_LOG_LEVEL")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("error "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_UPNP_DISABLE")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("true "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_PEER_BIND_ADDRESS")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token number"}},[a._v("0.0")]),a._v(".0.0:17002 "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_PEER_PRIVATE_KEY")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("bagacnag4afamw4hk6fk2yqr7fyug4cpf6fs6pzhkyttmwt4m27emmc76sdjcg6o7xfkyfrhw5gs3sx6r5aorj66xnveeo7wpeu456c4vsp3ga4z6ie "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_SONAR_PING_PEERS")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("bahwqdag4aeqcpluymxb3vby3xqr5py2ioamavw7vhas5n3ffvslhgbhcr5pehky "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token assign-left variable"}},[a._v("NIMONA_PEER_BOOTSTRAPS")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("bahwqdag4aeqllbusjc7e5qnsgq2e4z74g5kp7fsy53d2jin2bgwv5hel3wpbrty@tcps:0.0.0.0:17000 "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\ngo run ./cmd/sonar/main.go\n")])])])])}),[],!1,null,null,null);s.default=n.exports}}]);