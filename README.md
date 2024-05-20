# Opis Programu

## Opis

Ten program napisany za pomocą biblioteki gin implementuje prosty serwer HTTP, który obsługuje operacje na postach dotyczących ataków rekinów. Serwer umożliwia publikowanie, pobieranie oraz usuwanie postów. Dane o atakach rekinów są wczytywane z pliku JSON i przechowywane w pamięci serwera.

### Struktura Danych

Dane o atakach rekinów są reprezentowane przez strukturę `SharkAttack`, która zawiera następujące pola:
- `Date`: Data ataku
- `Country`: Kraj, w którym doszło do ataku
- `Name`: Imię ofiary
- `Activity`: Aktywność, którą wykonywała ofiara w momencie ataku
- `Age`: Wiek ofiary
- `Injury`: Obrazenia

### Funkcje Serwera

Serwer oferuje następujące funkcjonalności:

1. **Pobieranie wszystkich postów**:
   - Endpoint: `GET /posts`
   - Opis: Zwraca wszystkie zapisane posty o atakach rekinów.

2. **Dodawanie nowego posta**:
   - Endpoint: `POST /posts`
   - Opis: Dodaje nowy post na podstawie danych przesłanych w formacie JSON.
   - Przykładowe dane JSON:
     ```json
     {
       "date": "2024-05-04",
       "country": "PL",
       "name": "Jan Kowalski",
       "activity": "Fishing",
       "age": "47",
       "injury": "Fatal"
     }
     ```

3. **Pobieranie posta według ID**:
   - Endpoint: `GET /posts/:id`
   - Opis: Zwraca post o podanym ID.

4. **Usuwanie posta według ID**:
   - Endpoint: `DELETE /posts/:id`
   - Opis: Usuwa post o podanym ID.

