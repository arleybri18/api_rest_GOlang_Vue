
Rest API GO - Consume With App Vue
=====================

This is a project test implemented in GO language with  BD cockroach and using teh framework Vuejs.

## Getting Started

Make a copy of this project up and compile **C** files using *gcc* compiler see installing section.

### Prerequisites

You need the follow software to run the the program:

* GO version go1.12.5
* [cockroach](https://www.cockroachlabs.com)

### Configuration

Keep in mind the following recommendations:

1 - Port of listening backend server:

    Change in the file main.go in this variable
```
	port := ":5000"
```
2- Values to bd server:
    Change in the main.go and insertData.go files those global variables:
```
    var user string = "yonydb"
    var host_server string = "localhost"
    var port_server string = "26257"
    var name_bd string = "infodomains"
```

## App VUE js

The code of App web it's in the infoDomainsApp folder
    Change in the following files, the port of listening of the REST API app:: 

    infoDomainsApp\src\components\DomainInfo.vue

    ```
    var url = 'http://localhost:5000/domain/'+this.domain;
    ```

    infoDomainsApp\src\components\Report.vue

    ```
    var url = 'http://localhost:5000/report/';
    ```

## Warning

If you have a problem with CORS domain policy, install the next extension for google chrome. 
[Allow CORS: Access-Control-Allow-Origin](https://chrome.google.com/webstore/detail/allow-cors-access-control/lhobafahddgcelffkeicbaginigeejlf)

## Authors

* **Yony Briñez** - [Git Hub Repository](https://github.com/arleybri18/)