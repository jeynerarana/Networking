# HTTP Server

This is a simple http server Version 1.1. The server runs on port 8080 and will serve a webpage.

## Operating System

**Ubuntu:** 

So far the application only work with my Ubuntu machine and may not work work in another OS, the 
program was only tested on the Ubuntu System, but should work on another Operating Systems that are Linux based.


# How to Use:
**Running the Server**

1. Start the program by running `sudo go run <location of go file>` in my case I will use `sudo go run .` Since 
I only have **One** Go file in my current working directory. But if you have multiple go files please use `sudo go run nameOfFile.go`. You can leave the tab open and go to another tab and run the client.

**Running the Client**

2. To run the HTTP Client is easy. You can use your web browser. **Firefox preferably**
3. To connect go to the following URL `localhost:8080/` or you can simply put your private IP ie: `192.168.1.9:8080/`.
	1. Here is the File Structure:
<br />.<br />├── go.mod<br />├── main.go<br />└── www<br />&emsp;├── 404.html<br />&emsp;├──css<br />&emsp;│  ├── app.css<br />&emsp;│   ├── custom.css<br />&emsp;│   ├── foundation.css<br />&emsp;│   └── foundation.min.css<br />&emsp;├── img<br />&emsp;│   └── background.png<br />&emsp;├── index.html<br />&emsp;└── js<br />&emsp;&emsp;&emsp;├── app.js<br />&emsp;&emsp;&emsp;└── vendor<br />&emsp;&emsp;&emsp;&emsp;&emsp;├── foundation.js<br />&emsp;&emsp;&emsp;&emsp;&emsp;├── foundation.min.js<br />&emsp;&emsp;&emsp;&emsp;&emsp;├── jquery.js<br />&emsp;&emsp;&emsp;&emsp;&emsp;└── what-input.js

&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;5 directories, 14 files

4. Now you can use your browser to access these file by
requesting a file like so: `http://localhost:8080/404.html` <br /><br />Of course it won't be a good server if you can just get one file at a time. That's why you can can simply get the file `index.html` <br /><br />This files contains references to other files, and the browser will make multiple GET requests from the server to load the page. The requests will be handled by the server and correctly load the page.

**It can do HEAD requests too!!**

5. The Final functionality that the HTTP server has is to handle HEAD requests. 

	1. This part is simple, just run the server from **Running the Server** section and open a new Terminal.

	2. Now you need to run the following command `curl -X HEAD -I localhost:8080/` this will return just the headers from the server(No Body will be returned). You can can try calling directories that don't exist and you will get a 404 response.

# Expected Ouputs:
**GET Request**

1. The output for this should be all the Headers Including the Body which should be the data being sent. If you run a browser you would see the website if you run `index.html`, but you can also request the the `index.html` with the **curl** command as shown : `curl -X GET -i localhost:8080/index.html`, **Note: use lowercase 'i' for the command to get the body**
**Hint: when you run this in the browser you will get some guy in a bananna costume**

2. Expected Output for the command shown above:

		HTTP/1.1 200 OK
		Server: cihttpd
		Last-Modified: Fri Nov 12 22:54:46 2021
		Content-Length: 4096

		<!doctype html>
		<html class="no-js" lang="en" dir="ltr">
		    <head>
		        <meta charset="utf-8">
		        <meta http-equiv="x-ua-compatible" content="ie=edge">
		        <meta name="viewport" content="width=device-width, initial-scale=1.0">
		        <title>Kevin Scrivnor</title>
		        <link rel="stylesheet" href="css/foundation.css">
		        <link rel="stylesheet" href="css/app.css">
		        <link rel="stylesheet" href="css/custom.css">
		    </head>
		    <body>
		        <div class="top_padding"></div>
		        <div class="row">
		            <div class="large-12 columns">
		                <div class="callout">
		                    <h4>when you finally get the lab to work but also happen to be a banana</h4>
		                </div>
		            </div>
		        </div>

		        <script src="js/vendor/jquery.js"></script>
		        <script src="js/vendor/what-input.js"></script>
		        <script src="js/vendor/foundation.js"></script>
		        <script src="js/app.js"></script>}
		    </body>
		</html>


**HEAD Request**

1. The output for this should be just all the Headers.You can also request the the `index.html` with the **curl** command as shown here : `curl -X HEAD -I localhost:8080/index.html`, **Note: use uppercase 'I' in this command, both 'i' and 'I' work but 'i' will print out the bytes remaining to get on this page(if any)**

2. Expected Output for the command shown above(usinf `-I` as an argument):

		HTTP/1.1 200 OK
		Server: cihttpd
		Last-Modified: Fri Nov 12 22:54:46 2021
		Content-Length: 4096

3. Expected Output for the command shown above(usinf `-i` as an argument):

		
		Warning: Setting custom HTTP method to HEAD with -X/--request may not work the 
		Warning: way you want. Consider using -I/--head instead.
		HTTP/1.1 200 OK
		Server: cihttpd
		Last-Modified: Fri Nov 12 22:54:46 2021
		Content-Length: 4096
		curl: (18) transfer closed with 4096 bytes remaining to read





# Issues:
1. The browser used must be Firefox, I could not get it to work with others. In some cases if you run the server with chrome, half the image might show up. In some cases Google chrome would as for other files not specified. Or in other cases the browser won't load the page.

2. When I run the server I cannot get the `index.html` to show up whe the client requests `http://localhost:8080`. The `index.html`, would only be loaded when you request `http://localhost:8080/`(Includes '/') or calling `http://localhost:8080/index.html`