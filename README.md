# ai-smart-audit
A tool that takes a smart contract (e.g., Solidity code) as input and uses Azure AI services to: Detect potential vulnerabilities or gas inefficiencies (via NLP &amp; LLM), Summarise the smart contract logic in natural language, Optionally generate test cases or suggest improvements

## Folder Structure
``` bash
.
├── cmd
│   └── http
│       ├── main.go
│       └── server
│           └── server.go
├── config
│   └── config.yaml
├── deployment
├── docs
├── go.mod
├── internal
│   ├── adapters
│   │   ├── inbound
│   │   └── outbound
│   └── core
│       ├── models
│       │   └── dto
│       ├── port
│       │   ├── inbound
│       │   └── outbound
│       ├── services
│       └── utils
└── README.md
```
