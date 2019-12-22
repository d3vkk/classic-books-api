from page1 import page1
from page2 import page2
from page3 import page3
from page4 import page4

from bs4 import BeautifulSoup
import json
import re
from random import randrange

soup = BeautifulSoup(page3, 'html.parser')

count_array = []

# NUMBER OF BOOKS
list_body = soup.find(class_='list-body')
# print(listbody)
for count_column in list_body.find_all('div', {'class': 'col'}):
    for count_h4 in count_column.find_all('h4'):
        for count_title in count_h4.select('a')[0]:
            count_array.append(count_title)
# print(len(count_array))
num_of_books = len(count_array)

print('[')
# BOOK DETAILS
i = 0
while i < num_of_books:
    # print(i)
    print('{')
    print('"id":'+json.dumps(i)+',')
    # print('"id":'+json.dumps(i+151)+',')
    # isbn is a random number between 1 and 1million
    print('"isbn":'+json.dumps(randrange(1000000))+',')

    column = list_body.find_all('div', {'class': 'col'})[i]
    # print(column)

    h4 = column.find('h4')
    # print(h4)

    # BOOK TITLE
    book_title = h4.find_all('a')[0]
    # remove whitespace
    prep_title_values = re.sub('[\s+][\s+]', '', book_title.text)
    title_values = prep_title_values.strip()
    # print(title_values)
    print('"title":' + json.dumps(title_values) + ',')

    # BOOK AUTHOR
    book_author = h4.find_all('a')[1]
    # print(book_author.text)
    prep_author_values = re.sub('[\s+][\s+]', '', book_author.text)
    author_values = prep_author_values.strip()
    # print(author_values)
    print('"author":' + json.dumps(author_values) + ',')

    # BOOK SYNOPSIS
    pb3 = column.find('div', {'class': 'pb-3'})
    # print(pb3)
    thirddiv = pb3.find_all('div')[2]
    # print(thirddiv)
    book_synopsis = thirddiv.find('p')
    # print(book_synopsis.text)
    prep_synopsis_values = re.sub('[\s+][\s+]', '', book_synopsis.text)
    synopsis_values = prep_synopsis_values.strip()
    # print(synopsis_values)
    print('"synopsis":' + json.dumps(synopsis_values))  # don't add comma

    # don't add comma for the last book
    if i == (num_of_books - 1):
        print('}')
    else:
        print('},')

    i = i + 1

print(']')
