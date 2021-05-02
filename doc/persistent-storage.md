# Persistent Storage Design

Please refer to this [diagram](sqlite-db-diagram-v0.01.pdf) for more details

# SQLite db as a persistent storage 
  Persistent storage is needed to store data in a non-volatile device during and after the running of object store harvestor
  [SQLite](www.sqlite.org) has been chosen for that purpose.
  Object store harvestor by defualt stores all the data in db file : harvestor.db

# Object store harvestor DB persistent storage objectives
 - media files locations and types 
 - sidecar locations
 - partial response from media file upload API POST request
 - partial response from media file metadata API POST request 

# Desktop SQLite managing tool:
  - SQLite DB desktop tool : [sqlitebrowser](https://sqlitebrowser.org/)

# Working with SQLite db file : harvestor.db
  - Download and install [sqlitebrowser](https://sqlitebrowser.org/)
  - Open sqlitebrowser application
  - ![open sqlitebrowser](sqlitebrowser_open_db.png?raw=true "Using the menu navigate to your harvestor.db and open it")
  - Click on Browse Data
  - ![open sqlitebrowser](sqlitebrowser_browse_data.png?raw=true "Using the menu navigate to your harvestor.db and open it")
  - On the left hand side choose from the Table dropdown menu a table of your interest.
