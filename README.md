# CAKeyManager
this is a test project to demonstrate. a local distributed key manager

## Purpose
The purpose of this project is to demonstrate the concept of a distributed key manager
The key manager serves a storage to host secure keys 
In addition it serves as an X509 certificate authority for self signed keys and an SSH CA
The concepts that will be tested are:
* multiple nodes solution to provide ha / scale for the key store
* no single node stores a full key, only parts of it. The concept is based on SSS (shamir secret store) 
