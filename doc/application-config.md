# Design object-store-harvestor

Please refer to this [diagram](design-diagram-v0.02.pdf)

# Packages of object-store-harvester
  - Config (provides an ability to specify config parameters for all components below)
  - Orchestrater (Invokes File walker and http client)
  - File walker (recursively traverse directories and find files, connect to DB via DB client and store file paths in Local DB and mark them as 'new')
  - http client (Connects to DB via DB client, get file paths with mark 'new' and try to upload a file to API and mark the file ('success' or 'fail'))
  - DB client (provides connectivity to DB)
  - Local DB (persistence storage)
  - KeyCloak (OAuth 2 adapter for http client)
  - Common Logger (provides a way to log from each component above to a file)

# Deliverable components
    object-store-harvester binary executable file (Linux/Windows)
    object-store-harvester config file: harvestor_config.yml 

# Invoking object-store-harvester
- Linux
    `./binary_executable_file harvestor_config.yml` 

# Use cases
  - The Application runs on the Server (Phase 1)
  - The Application runs on the local PC (Future Phase)

# Harvestor Workflow (Phase 1)
  - Orchestrator initializes all dependencies
  - Orchestrator starts File walker process
  - File walker recursively traverse directories and find sidecar files
  - File walker reads and parse each sidecar file to identify original and derivative media files
  - File walker connects to DB via DB client and store file paths for original and derivative media files in Local DB and mark them as 'new'
  - File walker connects to DB via DB client and store sidecar file path in Local DB and mark them as 'new'
  - Orchestrator starts Http Client process
  - Http client connects to DB via DB client and upload all 'new' files one by one to related buckets ('original' & 'derivatives') and stores all responses to DB via DB client 
  - Http client posts metadata only for each 'original' uploaded file and stores all responses to DB via DB client
  - Http client via DB client pull all new sidecar files to build relationships: 
    - http client posts managed metadata relationship for each new sidecar based on the response data from DB from previous calls
    - http client posts derivative relationship for each new sidecar based on the response data from DB from previous calls

# An example with expectations from API server (Phase 1)
### Test set 
    cr2 file
    jpeg derivative
    yaml sidecar
### The sequence of operation is the following:
    Upload cr2 file /api/v1/file/{bucket}, record the returned id (file1)
    Upload jpeg file to /api/v1/file/{bucket}/derivative, record the returned id (file2)
    Create metadata resource (api/v1/metadata) and set fileIdentifier to file1, record the returned id (metadata1)
    Send managed attributes for metadata1
    Create derivative resource (/api/v1/derivative), set fileIdentifier to file2 and acDerivedFrom to metadata1
