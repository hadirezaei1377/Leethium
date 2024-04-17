Today I had a question about how to write a program with Golang to connect to a site that has tables and read the desired values from those tables.
For example, this site has a table

https://www.michael-smith-engineers.co.uk/resources/useful-info/approximate-viscosities-of-common-liquids-by-type

I want to go to the terminal for any material I entered and read its viscosity in this table and return it to me.
To write a program using Golang that connects to a website, reads information from that website's table and displays it to us, we can use libraries like "goquery" to process HTML and "net/http" to send HTTP requests. .

First, we need to design our application using "net/http" to send HTTP request to the website. Then use "goquery" to process the HTML and extract the required information from the page.

This program sends a GET request to the desired URL and then processes the information in the table using the "goquery" library. You can then get the viscosity of any liquid you've entered into the program using the getViscosity function.