# Exercise #8: Phone Number Normalizer

## Exercise details

This exercise is fairly straight-forward - we are going to be writing a program that will iterate through a database and normalize all of the phone numbers in the DB. After normalizing all of the data we might find that there are duplicates, so we will then remove those duplicates keeping just one entry in our database.

The primary goal is to create a program that normalizes the database.

We will also need to explore some basic string manipulation techniques to normalize our phone numbers

On to the exercise - we will start by creating a database along with a `phone_numbers` table. Inside that table we want to add the following entries (yes, I know there are duplicates):

```
1234567890
123 456 7891
(123) 456 7892
(123) 456-7893
123-456-7894
123-456-7890
1234567892
(123)456-7892
```

You can organize your table however you want, and you may add whatever extra fields you want. 

Once you have the entries created, our next step is to learn how to iterate over entries in the database using Go code. With this we should be able to retrieve every number so we can start normalizing its contents.

Once you have all the data in the DB, our next step is to normalize the phone number. We are going to update all of our numbers so that they match the format:

```
##########
```

That is, we are going to remove all formatting and only store the digits. When we want to display numbers later we can always format them, but for now we only need the digits.

In the list of numbers provided, the first entry, along with the second to last entry, match this format. All of the others do not and will need to be reformatted. There are also some duplicates that will show up once we have reformatted all the numbers, and those will need removed form the database but don't worry about that for now.

Once you written code that will successfully take a number in with any format and return the same number in the proper format we are going to use an `UPDATE` to alter the entries in the database. If the value we are inserting into our database already exists (it is a duplicate), we will instead be deleting the original entry.

When your program is done your database entries should look like this (the order is irrelevant, but duplicates should be removed):

```
1234567890
1234567891
1234567892
1234567893
1234567894
```

