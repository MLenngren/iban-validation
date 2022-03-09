<div align="center">
  <img src="https://raw.githubusercontent.com/MLenngren/iban-validator/main/ibaninfo.jpg" alt="Iban info" width="400" />
</div>

<div align="center">
  <a href="https://twitter.com/golang">
    <img src="https://badgen.net/twitter/follow/golang" alt="go golang twitter" />
  </a>
  <a href="#">
    <img src="https://camo.githubusercontent.com/abb97269de2982c379cbc128bba93ba724d8822bfbe082737772bd4feb59cb54/68747470733a2f2f63646e2e7261776769742e636f6d2f73696e647265736f726875732f617765736f6d652f643733303566333864323966656437386661383536353265336136336531353464643865383832392f6d656469612f62616467652e737667" alt="awesome" />
  </a>
</div>


# IBAN Validator
Validates an IBAN that is in the IBAN register at 3.1 below, as such any experimental IBAN's do not validate since we do not have those in the ibanRegisterInfo list which is used to verify the length of the IBAN.


## Tasks: 

1. Understand how to validate an IBAN
2. Find a way to get mod-97 (as described in ISO 7064) in go, to be able to convert and validate a large number

And a requirement is that it should not be an IBAN library. After a quick search most of the packages/libraries is IBAN related.

3. Get an up to date list of the IBAN Registry (optional, but helpful)

4. Create function to validate IBAN number given the information in 1.1 below.

5. Setup docker

## Info for tasks:

1. https://en.wikipedia.org/wiki/International_Bank_Account_Number
1.1. Excerpt from Wiki
Algorithms

Validating the IBAN
An IBAN is validated by converting it into an integer and performing a basic mod-97 operation (as described in ISO 7064) on it. If the IBAN is valid, the remainder equals 1.[Note 1] The algorithm of IBAN validation is as follows:[8]

Check that the total IBAN length is correct as per the country. If not, the IBAN is invalid
Move the four initial characters to the end of the string
Replace each letter in the string with two digits, thereby expanding the string, where A = 10, B = 11, ..., Z = 35
Interpret the string as a decimal integer and compute the remainder of that number on division by 97
If the remainder is 1, the check digit test is passed and the IBAN might be valid.

Example (fictitious United Kingdom bank, sort code 12-34-56, account number 98765432):

• IBAN:		GB82 WEST 1234 5698 7654 32	
• Rearrange:		W E S T12345698765432 G B82	
• Convert to integer:		3214282912345698765432161182	
• Compute remainder:		3214282912345698765432161182	mod 97 = 1
[8] https://www.ecbs.org/Download/EBS204_V3.2.pdf

3. https://www.swift.com/search?keywords=IBAN+list&search-origin=result_search
Searched on Swift, they had a txt file and a pdf file of IBAN registers. Download and look at the country details, automate download and parsing to always have an up to date IBAN country list. For now I copied all the text and recorded a macro(notepad++) and created the array with country codes and IBAN length's. So it's easy to update that way. 

3.1. TODO:  Automatic update the IBAN register
- Download https://www.swift.com/swift-resource/9606/download
- Parse out the TEXT 
- Insert into DB


## NOTICE

- This is my first go


### PONDERINGS

- What is some of the best microservice structures in go
- How to log best practice
- As above for errors, or we take it its the same sort of since go dont have exception handling like I am used ot.


### some commands

- go get github.com/stretchr/testify