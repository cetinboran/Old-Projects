# Basic Web Security Checker
+ This project is a simple one that allows you to search for XSS and SQL injection on the endpoints of the websites you desire.

# Installation
+ You can download the project by using the following command: `git clone https://github.com/cetinboran/BasicSEC`

# Database
+ To use the project, you should import the .sql file included in the project into MySQL.

# How to use?
+ Download and run the project.
+ Sign up.
+ Add the URL you want through the "add Url" tab.
+ Press the "scan" button in the URL field that appears on the screen
+ After entering the necessary scan options, press the "scan" button. It will automatically initiate the scan and add the results to the database.
+ It initially sends an empty request to the website, and it adds to the database the entries with different content lengths and statuses.
+ Once the scanning process is complete, you can view the scans conducted for a particular URL by clicking on the "view" option located under the URL section on the home page.

# You have to know.
+ In scan page params format is username:* password:root each parameter is written on a single line. The '*' is the keyword for place the lines here.

# Db Configs
+ You have to add the .sql file to your mysql workbench.
+ Create your db.env file in /. You can see an example below
  + **DB_USER** = "root"
  + **DB_PASS** = "root."
  + **DB_HOST** = "localhost"
  + **DB_PORT** = "3306"
  + **DB_NAME** = "basic_security"

# Contact
<p align="center">
  <a href="https://github.com/cetinboran">
    <img src="https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/github.svg" alt="github" height="40">
  </a>
  <a href="https://www.linkedin.com/in/cetinboran-mesum/">
    <img src="https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/linkedin.svg" alt="linkedin" height="40">
  </a>
  <a href="https://www.instagram.com/2023an_m/">
    <img src="https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/instagram.svg" alt="instagram" height="40">
  </a>
  <a href="https://twitter.com/2023anM">
    <img src="https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/twitter.svg" alt="twitter" height="40">
  </a>
</p>
