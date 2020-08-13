package api

// func (api *API) HandleGetStreams(c *router.Context) {
// 	write := func(conn *websocket.Conn, data interface{}) error {
// 		return conn.WriteJSON(data)
// 	}

// 	wsupgrader := websocket.Upgrader{
// 		ReadBufferSize:  1024,
// 		WriteBufferSize: 1024,
// 		CheckOrigin: func(r *http.Request) bool {
// 			return true
// 		},
// 	}

// 	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		c.AbortWithError(500, err) // nolint: errcheck
// 		return
// 	}

// 	ctx := context.Background()
// 	logger := log.FromContext(ctx).Named("api")
// 	incoming := make(chan object.Object, 100)
// 	outgoing := make(chan object.Object, 100)

// 	go func() {
// 		for {
// 			select {
// 			case v := <-incoming:
// 				m := api.mapObject(v)
// 				if err := write(conn, m); err != nil {
// 					// TODO handle error
// 					continue
// 				}

// 			case req := <-outgoing:
// 				sig, err := object.NewSignature(
// 					api.keychain.GetPrimaryPeerKey(),
// 					req,
// 				)
// 				if err != nil {
// 					logger.Error(
// 						"could not sign outgoing object",
// 						log.Error(err),
// 					)
// 					req.Set("_status", "error signing object")
// 					if err := write(conn, api.mapObject(req)); err != nil {
// 						// TODO handle error
// 						continue
// 					}
// 				}
// 				req = req.AddSignature(sig)
// 				// TODO(geoah) better way to require recipients?
// 				// TODO(geoah) helper function for getting subjects
// 				subjects := []string{}
// 				if ps := req.Get("_recipients"); ps != nil {
// 					if subsi, ok := ps.([]interface{}); ok {
// 						for _, subi := range subsi {
// 							if sub, ok := subi.(string); ok {
// 								subjects = append(subjects, sub)
// 							}
// 						}
// 					}
// 				}
// 				if len(subjects) == 0 {
// 					// TODO handle error
// 					req.Set("_status", "no subjects")
// 					// nolint: staticcheck
// 					if err := write(conn, api.mapObject(req)); err != nil {
// 						// TODO handle error
// 					}
// 					continue
// 				}
// 				for _, recipient := range subjects {
// 					rec := peer.LookupByOwner(crypto.PublicKey(recipient))
// 					if err := api.exchange.Send(ctx, req, rec); err != nil {
// 						logger.Error("could not send outgoing object",
// 							log.Error(err),
// 						)
// 						req.Set("_status", "error sending object")
// 					}
// 					// TODO handle error
// 					if err := write(conn, api.mapObject(req)); err != nil {
// 						// TODO handle error
// 						continue
// 					}
// 				}
// 			}
// 		}
// 	}()

// 	for {
// 		_, msg, err := conn.ReadMessage()
// 		if err != nil {
// 			if err == io.EOF {
// 				logger.Debug("ws conn is dead", log.Error(err))
// 				return
// 			}

// 			if websocket.IsCloseError(
// 				err,
// 				websocket.CloseGoingAway,
// 				websocket.CloseNormalClosure,
// 			) {
// 				logger.Debug("ws conn closed", log.Error(err))
// 				return
// 			}

// 			if websocket.IsUnexpectedCloseError(err) {
// 				logger.Warn(
// 					"ws conn closed with unexpected error",
// 					log.Error(err),
// 				)
// 				return
// 			}

// 			logger.Warn("could not read from ws", log.Error(err))
// 			continue
// 		}
// 		m := map[string]interface{}{}
// 		if err := json.Unmarshal(msg, &m); err != nil {
// 			logger.Error("could not unmarshal outgoing object", log.Error(err))
// 			continue
// 		}
// 		o := object.FromMap(m)
// 		outgoing <- o
// 	}
// }