# Go SDK for Salesforce Marketing Cloud

Used `gowsdl` to generate code from the Salesforce Marketing Cloud SOAP API WSDL.

```shell
$ gowsdl -o gen.go -p gen etframework.wsdl
```

## Work in progress

### SOAP API

- [ ] Data Folders
  - [ ] Create
  - [ ] Retrieve
  - [ ] Delete
- [ ] Data Extensions
  - [ ] Create
  - [ ] Retrieve
  - [ ] Delete
  - [ ] Clear
  - [ ] Add Records
  - [ ] Update Records
- [ ] Query Definitions
  - [ ] Create
  - [ ] Perform
  - [ ] Delete
  - [ ] Clear
- [ ] Import Definitions
  - [ ] Create
  - [ ] Perform
  - [ ] Retrieve Results Summary
  - [ ] Delete
- [ ] Email Send Definitons
  - [ ] Create
  - [ ] Retrieve
  - [ ] Update
  - [ ] Delete
  - [ ] Send
  - [ ] Test Send
  - [ ] Schedule Send
  - [ ] Cancel Send
  - [ ] Pause Send
  - [ ] Resume Send
- [ ] Extract Tracking Data
  - [ ] sends
  - [ ] opens
  - [ ] clicks
  - [ ] conversions
  - [ ] unsubs
  - [ ] bounces
  - [ ] spam complaints

### REST API

- [ ] Content Builder
- [ ] Send Preview Email
- [ ] SMS
