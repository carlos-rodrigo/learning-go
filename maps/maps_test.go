package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("know word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("unknow word", func(t *testing.T) {
		_, got := dictionary.Search("other")

		assertError(t, got, ErrNotFound)
	})

	t.Run("add new word", func(t *testing.T) {
		dictionary = Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("add a word that already exist", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExist)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("update definition", func(t *testing.T) {
		word := "test"
		definition := "this is a test definition"
		dictionary := Dictionary{word: definition}
		newDefinition := "this is a new test definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("update not existing word definitions should be fail", func(t *testing.T) {
		word := "test"
		newDefinition := "newDefinition"
		dictionary := Dictionary{}

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, ErrWordDoesNotExist)
	})

	t.Run("delete a word from dictionary", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)

		if err != ErrNotFound {
			t.Errorf("Expect an error because the word was deleted")
		}

	})
}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != definition {
		t.Errorf("got %q want %q", got, definition)
	}
}
