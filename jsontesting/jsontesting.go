package jsontesting

// BookData data in JSON format
func BookData() string {
	return `
    {
      "classicbooks": [
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
        }
      ]
    }
  `
}
