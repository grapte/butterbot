# ButterBot

This bot currently pulls from WhaleAlert APIs and then posts to a discord channel. Whales refer to large transactions of crypto currencies.
## File Descriptions

```
.
├── LICENSE
├── api
│   ├── ... // hand written openapi/cue/others... formats for whalealert APIs
├── butterbot.go  // main entry point for the bot
├── cmd
│   └── cmd.go              // command parsing and configuration reading
├── go.mod      // go modules user readable
├── go.sum      // cache file for modules
├── readme.md   // You are here
├── service
│   ├── discord
│   │   └── discord.go // discord service and all discord related code contexts
│   └── whalealert
│       ├── whalealert.go   // whalealert api pulling and request creation
│       └── whalealert_test.go // tests the requests
├── types
│   └── whalealertapi.go    // golang struct with json annotations for whalealert api
└── verify.cue                    // cue file for verifying secrets validity

```