easymarkdown
============

It's an application to automatically publish your local modified markdown document to the blog.

It includes two parts

Server part:
It's running on the server side. 

- It provide service on an port for web
- It provide service on an port for receiving markdown document upload. ( It need use /password)
- It maintains a folder including the markdown documenents
- It transforms the markdown format into html 


Client part:
It's running on the client side.

- It is located at the folder including the markdown document
- It has a configure file including user/password used to upload the markdown docuemnts.
- It can check what documents are newest and need to be upload.
- It can connect to server side and upload the newest markdown files
