from pickle import NONE
from dbconnection import db_connection
from urllib import request
from bs4 import BeautifulSoup
import mysql.connector
import requests



#import ipdb;ipdb.set_trace()

# Establishing db connection
def db_insert(productName,res,providerId):
    db=db_connection()
    mycursor=db.cursor()
    try: 
        print("INSERT INTO scraperdb.products(`name`,`itemName`,`rating`,`cost`,`noOfRatings`,`providerId`,`created_at`) VALUES(\"{}\",\"{}\",\"{}\",{},{},{},CURRENT_TIMESTAMP)".format(productName,res[0],res[1],res[2],res[3],providerId))
        mycursor.execute("INSERT INTO scraperdb.products(`name`,`itemName`,`rating`,`cost`,`noOfRatings`,`providerId`,`created_at`) VALUES(\"{}\",\"{}\",\"{}\",{},{},{},CURRENT_TIMESTAMP)".format(productName,res[0],res[1],res[2],res[3],providerId))
        db.commit() 
        print("db query executed successfully")
    except:  
        print("invalid insert details") 
        db.rollback()  
    db.close()

#Reviews
#1.keep configuration seperate from main code- just import mydb(write a module for that)-DRY
#2.Never use passwords in main file use .env file for sensitive information
def db_serach(productName):
    db=db_connection()
    mycursor=db.cursor()
    try: 
        mycursor.execute("SELECT * FROM scraperdb.products where name=\"{}\"".format(productName))
        records=mycursor.fetchall()
    except:  
        print("No such product to search!") 
        db.rollback()  
    db.close()
    return records

def db_update(res,productName):
    db=db_connection()
    mycursor=db.cursor()
    print("cursor created!")
    print(res[0])
    print(res[1])
    print(res[2])
    print(res[3])
    #import ipdb;ipdb.set_trace();
    try: 
        print("UPDATE `scraperdb`.`products` SET `itemName` = '{}', `rating` = '{}', `cost` = '{}', `noOfRatings` = '{}' WHERE (`name` = '{}')".format(res[0],res[1],res[2],res[3],productName))   
        mycursor.execute("UPDATE `scraperdb`.`products` SET `itemName` = '{}', `rating` = '{}', `cost` = '{}', `noOfRatings` = '{}' WHERE (`name` = '{}')".format(res[0],res[1],res[2],res[3],productName))
        db.commit() 
        print("db query executed successfully")
    except:  
        print("invalid search details") 
        db.rollback()  
    db.close()
    

def scrape_amazon_data(productName):
    amazon_url="https://www.amazon.in/s?k={}".format(productName)
    print(amazon_url)
    
    req=requests.get(amazon_url)
    soup=BeautifulSoup(req.content, "lxml")
    #print(req.text)
    
    try:
        #itemName
        name_tags=soup.find('div',class_="a-section a-spacing-none puis-padding-right-small s-title-instructions-style")
        name=name_tags.h2.text.split("-")[0]

        #ratings
        rating_tags=soup.find('i',class_="a-icon a-icon-star-small a-star-small-4-5 aok-align-bottom")
        rating=rating_tags.text.split(" ")[0]

        #cost
        cost_tags=soup.find('span',class_="a-price-whole")
        cost_temp=cost_tags.text.split(",")
        cost=""
        for i in cost_temp:
            cost+=i

        #number of ratings
        noOfRatings_tag=soup.find('div',class_="a-section a-spacing-none a-spacing-top-micro")
        noOfRating_temp=noOfRatings_tag.text.split(" ")[5].split(",")
        noOfRating=""
        for i in noOfRating_temp:
            noOfRating+=i

        print(name,rating,cost,noOfRating)
        return name,rating,cost,noOfRating
    except:
        print("Unable to receive data,Amazon server busy!")
    

def scrape_flipkart_data(productName):
    flipKart_url="https://www.flipkart.com/search?q={}".format(productName)
    
    req=requests.get(flipKart_url)
    soup=BeautifulSoup(req.content, "lxml")
    result=[]

    try:
        #itemName
        name_tags=soup.find('div',class_="_4rR01T")
        raw_name=name_tags.text
        name=""
        for i in raw_name:
            if i=="/" or i=="\"":continue
            name+=i
        result.append(name)

        #ratings
        rating_tags=soup.find('div',class_="_3LWZlK")
        rating=rating_tags.text

        result.append(rating)

        #cost
        cost_tags=soup.find('div',class_="_30jeq3 _1_WHN1")
        cost_tags+=""
        cost=""
        print(cost_tags)
        for data in cost_tags:
            if data=="₹" or data==",":continue
            cost+=data
        
        result.append(cost)

        #number of ratings
        noOfRatings_tag=soup.find('span',class_="_2_R_DZ")
        print("noOfRatings_tag",noOfRatings_tag)
        noOfRatings_temp=noOfRatings_tag.text.split(" ")[0].split(",")
        print(noOfRatings_temp)
        noOfRating=""
        for data in noOfRatings_temp:
            noOfRating+=data
        
        result.append(noOfRating)
        print(name,rating,cost,noOfRating)
        return name,rating,cost,noOfRating
        # return result
    except:
        print("Unable to receive data,FlipKart server busy!")

#print error object
    

if __name__=='__main__':
    # amazonData=scrape_flipkart_data("iphone") 
    # amazonData=scrape_amazon_data("iphone")
    # print(list(amazonData))
    # if amazonData:
    #     for name,rating,cost,noOfRatings in list(amazonData):
            
    # else:
    #     print("No data")

    #import ipdb;ipdb.set_trace()
       #change fn name 
    # if amazonData:
    #     name,rating,cost,noOfRatings =amazonData
    #     print("\nProduct Name :{}\nrating :{}\ncost :{}\nnoOfRatings :{}".format(name,rating,cost,noOfRatings))
    #     # db_insert("laptop",amazonData,1)


     
    userExists=True #better name for condition Say userOptedToExit 
    while(userExists):
        
        print("\nDo you want to search for a product?")
        print("1.Yes")
        print("2.No")
        userChoice=input("enter your choice:") #string comparision also will do  2.better name

        if(userChoice=='1'):
            #seach for data
            print("\nPlease Select Provider")
            print("1.Amazon")
            print("2.FlipKart")
            print("3.Both")
            userChoice=input("enter your choice:")
            productName=input("enter Product Name to search:")

            #have to change the logic

            if(userChoice=='1'):
                data=db_serach(productName)            #print searched data
                if data:
                    for _,_,itemName,rating,cost,noOfRatings,_,_ in data:
                        print("\nProduct Name :{}\nrating :{}\ncost :{}\nnoOfRatings :{}".format(itemName,rating,cost,noOfRatings))
                else:
                    amazonData=scrape_amazon_data(productName)    #change fn name 
                    if amazonData:
                        name,rating,cost,noOfRatings =amazonData
                        print("\nProduct Name :{}\nrating :{}\ncost :{}\nnoOfRatings :{}".format(name,rating,cost,noOfRatings))
                        db_insert(productName,amazonData,0)
                    # if amazonData:
                    #     db_insert(productName,amazonData,0)
                    #     for name,rating,cost,noOfRatings in amazonData:
                    #         print("\nProduct Name :{}\nrating :{}\ncost :{}\nnoOfRatings :{}".format(itemName,rating,cost,noOfRatings))

                #else print the data
                #same print logic for db_search and Amazonresult line no 162
                print("\nPlease Select your option")
                print("1.Refetch the data")
                print("2.Go back")
                userChoice=input("enter your choice:")
                if(userChoice=='1'):
                    amazonData=scrape_amazon_data(productName)
                    if amazonData:
                        name,rating,cost,noOfRatings =amazonData
                        print("\nProduct Name :{}\nrating :{}\ncost :{}\nnoOfRatings :{}".format(name,rating,cost,noOfRatings))
    #variables names in camelCase
                    print("\nPlease Select your option")
                    print("1.Update the data")
                    print("2.Go back")
                    userChoice=input("enter your choice:")

                    if(userChoice=='1' and amazonData):
                            db_update(amazonData,productName)
                            print("updation successful!")
                    elif userChoice=='2':
                        continue

                elif userChoice=='2':
                        continue


            elif(userChoice=='2'):
                data=db_serach(productName)            #print searched data
                for i in data:
                    print(i)
                if data is not None:
                        flipkartData=scrape_flipkart_data(productName)    #change fn name
                        
                        if flipkartData:
                            name,rating,cost,noOfRatings =flipkartData
                            print("\nProduct Name :{}\nrating :{}\ncost :{}\nnoOfRatings :{}".format(name,rating,cost,noOfRatings))
                            db_insert(productName,flipkartData,1)

                #else print the data
                #same print logic for db_search and Amazonresult line no 162
                print("\nPlease Select your option")
                print("1.Refetch the data")
                print("2.Go back")
                userChoice=input("enter your choice:")
                if(userChoice=='1'):
                    flipkartData=scrape_flipkart_data(productName)
                    if flipkartData is not None:
                        for i in flipkartData:
                            print(i)
    #variables names in camelCase
                    print("\nPlease Select your option")
                    print("1.Update the data")
                    print("2.Go back")
                    userChoice=input("enter your choice:")
                    if(userChoice=='1' and flipkartData):
                            db_update(flipkartData,productName)
                            print("updation successful!")

                    elif userChoice=='2':
                        continue

                elif userChoice=='2':
                        continue



            elif userChoice=='3':
                
                data=db_serach(productName)
                if data:
                    for i in data:
                        print(i)
                else:
                    amazonData=scrape_amazon_data(productName)
                    flipkartData=scrape_flipkart_data(productName) 
                    
                    db_insert(productName,amazonData,0)
                    db_insert(productName,flipkartData,1)
                    
                    if amazonData:
                        name,rating,cost,noOfRatings =amazonData
                        print("\nProduct Name :{}\nrating :{}\ncost :{}\nnoOfRatings :{}".format(name,rating,cost,noOfRatings))
                    if flipkartData:
                        name,rating,cost,noOfRatings =flipkartData
                        print("\nProduct Name :{}\nrating :{}\ncost :{}\nnoOfRatings :{}".format(name,rating,cost,noOfRatings))

                print("\nPlease Select your option")
                print("1.Refetch the data")
                print("2.Go back")
                userChoice=input("enter your choice:")
                if(userChoice=='1'):
                    amazonData=scrape_amazon_data(productName)
                    flipkartData=scrape_flipkart_data(productName) 
                    if amazonData:
                        name,rating,cost,noOfRatings =amazonData
                        print("\nProduct Name :{}\nrating :{}\ncost :{}\nnoOfRatings :{}".format(name,rating,cost,noOfRatings))
                    if flipkartData:
                        name,rating,cost,noOfRatings =flipkartData
                        print("\nProduct Name :{}\nrating :{}\ncost :{}\nnoOfRatings :{}".format(name,rating,cost,noOfRatings))

                    print("\nPlease Select your option")
                    print("1.Update the data")
                    print("2.Go back")
                    userChoice=input("enter your choice:")
                    if(userChoice=='1'):
                        db_update(amazonData,productName)
                        db_update(flipkartData,productName)
                        print("db updation successful!")

                    elif userChoice=='2':
                        continue

                elif userChoice=='2':
                        continue

        elif(userChoice=='2'):
            userExists=False







#Review
#instead of if and elif use dictionary
#print db data by formatting it prints statements can rather be wrapped them in a function

    #print("result is:",result)
    # res=list(map(lambda i,j,k:(i,j,k),n,r,c))
    #db_connection(result)

    

    # BACKUP
    #  #itemName
    # name_tags=soup.find('span',class_="a-size-medium a-color-base a-text-normal")
    # print(name_tags.text)
    # # name=name_tags.h2.text.split("-")[0]
    # # print(names)
    
    # #ratings
    # rating_tags=soup.find('div',class_="a-section a-spacing-none a-spacing-top-micro")
    # if(rating_tags.a!=None):
    #     print(rating_tags.a.text[0])
    #     rating=rating_tags.a.text[0]
    
    # #cost
    # cost_tags=soup.find('span',class_="a-price-whole")
    # print(cost_tags.text)
    # cost=cost_tags.text
    # #number of ratings
    # noOfRatings_tag=soup.find('span',class_="a-size-base puis-light-weight-text s-link-centralized-style")
    # print(noOfRatings_tag.text)






        
    

    # #itemName
    # name_tags=soup.find('div',class_="a-section a-spacing-none puis-padding-right-small s-title-instructions-style")
    # print(name_tags.h2.text.split("-")[0])
    # name=name_tags.h2.text.split("-")[0]
    # # print(names)
    
    # #ratings
    # rating_tags=soup.find('div',class_="a-section a-spacing-none a-spacing-top-micro")
    # if(rating_tags.a!=None):
    #     print(rating_tags.a.text[0])
    #     rating=rating_tags.a.text[0]
    
    # #cost
    # cost_tags=soup.find('span',class_="a-price-whole")
    # print(cost_tags.text)
    # cost=cost_tags.text
    # #number of ratings
    # noOfRatings_tag=soup.find('span',class_="a-size-base puis-light-weight-text")
    # print(noOfRatings_tag.text)
    # # noOfRatings=noOfRatings_tag.text



    #  UPDATE `scraperdb`.`products` SET `itemName` = \'{}\', `rating` = \'{}\', `cost` = \'{}\', `noOfRatings` = \'{}\' WHERE (`name` = `{}``)
        #mycursor.execute("UPDATE `scraperdb`.`products` SET `itemName` = \'{}\', `rating` = \'{}\', `cost` = \'{}\', `noOfRatings` = \'{}\' WHERE (`name` = \'{}\'))".format(res[0],res[1],res[2],res[3],inputString))
        #mycursor.execute("UPDATE `scraperdb`.`products` SET `itemName` = '{}', `rating` = '{}', `cost` = '{}', `noOfRatings` = '{}' WHERE (`name` = '{}')".format(res[0],res[1],res[2],res[3],inputString))
        # records=mycursor.fetchall()
        # print("Database Result:")
        # for i in records:
        #     print(i)