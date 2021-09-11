# Classic Books API

![Classic Books](https://github.com/d3vkk/classic-books-api/blob/master/classic-books.png)

Get Classic Books Data. Built with with Golang. Data scraped with Python

## Methods

### Get All Classic Books

URL

[`http://localhost:3000/api/books`](http://localhost:3000/api/books)

Response

```json
[
  {
    "id": 0,
    "isbn": 60684,
    "title": "In Search of Lost Time",
    "author": "Marcel Proust",
    "synopsis": "Swann's Way, the first part of A la recherche de temps perdu, Marcel Proust's seven-part cycle, was published in 1913. In it, Proust introduces the themes that run through the entire work. The narr..."
  },
  {
    "id": 1,
    "isbn": 688150,
    "title": "Ulysses",
    "author": "James Joyce",
    "synopsis": "Ulysses chronicles the passage of Leopold Bloom through Dublin during an ordinary day, June 16, 1904. The title parallels and alludes to Odysseus (Latinised into Ulysses), the hero of Homer's Odyss..."
  },
  {
    "id": 2,
    "isbn": 963600,
    "title": "Don Quixote",
    "author": "Miguel de Cervantes",
    "synopsis": "Alonso Quixano, a retired country gentleman in his fifties, lives in an unnamed section of La Mancha with his niece and a housekeeper. He has become obsessed with books of chivalry, and believes th..."
  }
  ...
]
```
### Classic Book Search By Id

URL

[`http://localhost:3000/api/books/174`](http://localhost:3000/api/books/174)

Response

```json
{
  "id": 174,
  "isbn": 108452,
  "title": "Treasure Island",
  "author": "Robert Louis Stevenson",
  "synopsis": "Traditionally considered a coming-of-age story, it is an adventure tale known for its superb atmosphere, character and action, and also a wry commentary on the ambiguity of morality\u2014as seen in Long..."
}
```

Fork or clone this repo
```
git clone https://github.com/d3vkk/classic-books-api.git
```

© 2019-present Donald K • Under MIT License
