1. Czym jest transakcja i czym się różni od procedury
Transakcja to spójna jednostka pracy w bazie danych, która musi spełniać warunki ACID (atomowość, spójność, izolacja, trwałość). Procedura to zapisany zbiór instrukcji SQL. Różnica: transakcja gwarantuje bezpieczeństwo danych, a procedura to tylko program w bazie – może zawierać wiele transakcji.
2. Metody modelowania procesów biznesowych
Najczęściej stosuje się: BPMN (notacja graficzna), UML (diagramy aktywności, przypadków użycia), EPC (łańcuchy procesów zdarzeń), IDEF0. Służą one do analizy i optymalizacji procesów w organizacji.
3. Metody badania złożoności algorytmów (O i o)
Złożoność bada się przez analizę liczby operacji w zależności od wielkości danych. Duże O (O) określa górne ograniczenie, a małe o (o) oznacza ścisłą górną granicę. Np. O(n²) oznacza, że czas rośnie maksymalnie kwadratowo. Ważne są klasy: logarytmiczna, liniowa, kwadratowa, wykładnicza.
4. Co to są obiekty pozorne i do czego mogą być zastosowane (Java)
Obiekt pozorny (mock, stub) to sztuczna implementacja klasy lub interfejsu używana w testach. Dzięki nim można testować fragment kodu bez potrzeby korzystania z prawdziwych zasobów (np. bazy danych).
5. Paradygmat programowania obiektowego
Opiera się na czterech filarach: enkapsulacja (ukrywanie szczegółów), dziedziczenie (przekazywanie cech), polimorfizm (różne zachowanie tej samej metody), abstrakcja (uproszczenie złożoności). Celem jest łatwiejsze modelowanie rzeczywistości.
6. Sieci neuronowe – pojęcie i zastosowania
To modele matematyczne inspirowane pracą mózgu, składające się z neuronów i warstw. Stosowane do rozpoznawania obrazów, przetwarzania języka naturalnego, prognozowania i klasyfikacji.
7. Postaci normalne relacji
1NF – brak powtarzających się grup danych, atomowość.
2NF – spełnia 1NF i brak częściowej zależności od klucza.
3NF – spełnia 2NF i brak zależności przechodnich.
Celem normalizacji jest eliminacja redundancji.
8. Klasteryzacja (grupowanie) danych
Polega na podziale obiektów na grupy (klastry), gdzie obiekty w klastrze są do siebie podobne, a różne między sobą. Przykład: algorytm k-średnich, hierarchiczne, DBSCAN.
9. Klasyfikacja języków programowania
imperatywne (C, Pascal),
deklaratywne (SQL, Prolog),
obiektowe (Java, C++),
proceduralne (C).
funkcyjne (Haskell, Scala),
logiczne (Prolog),
skryptowe (Python, JavaScript).
10. Numeryczna metoda wyznaczania całki oznaczonej
Najczęściej stosuje się: metodę prostokątów, trapezów i Simpsona. Polega na przybliżaniu pola pod krzywą za pomocą prostych figur geometrycznych.
11. Inteligentne techniki obliczeniowe
To metody inspirowane naturą, np. sztuczne sieci neuronowe, algorytmy genetyczne, logika rozmyta. Stosowane w AI i optymalizacji.
12. Warunki dla relacji w 3NF
Relacja musi być w 2NF i każdy atrybut niekluczowy musi zależeć tylko od klucza głównego, a nie od innych atrybutów.
13. Co to jest klasa
To szablon definiujący strukturę (pola, właściwości) i zachowanie (metody) obiektów. Na podstawie klasy tworzy się obiekty.
14. Transakcje w systemach współbieżnych
Aby unikać błędów (np. dirty read, phantom read), stosuje się poziomy izolacji: Read Uncommitted, Read Committed, Repeatable Read, Serializable. Do koordynacji wielu baz używa się protokołu dwufazowego (2PC).
15. Mechanizm dziedziczenia w programowaniu obiektowym
Pozwala tworzyć nowe klasy na bazie już istniejących. Ułatwia ponowne użycie kodu i tworzenie hierarchii.
16. Proces uczenia sieci neuronowej
Polega na dostosowywaniu wag połączeń między neuronami, najczęściej metodą propagacji wstecznej błędu. Celem jest minimalizacja błędu prognozy.
17. Algorytm tyrana wyboru nowego koordynatora procesów
Najwyższy numer procesu przejmuje rolę koordynatora. Procesy wysyłają wiadomości do „silniejszych” procesów, aż zostanie wybrany jeden zwycięzca.
18. Mechanizm działania wskaźnika w programowaniu
Wskaźnik to zmienna przechowująca adres w pamięci. Umożliwia dostęp do danych i manipulację nimi w sposób pośredni.
19. Problem NP-zupełny
To zadanie, dla którego nie znamy algorytmu rozwiązującego w czasie wielomianowym, a każde inne NP-problem można do niego sprowadzić. Przykład: problem komiwojażera.
21. System operacyjny – pojęcie i budowa
To oprogramowanie pośredniczące między sprzętem a aplikacjami. Składa się z jądra, modułu zarządzania pamięcią, procesami, plikami i urządzeniami.
22. Pojęcie odległości i przykłady
Odległość mierzy podobieństwo/dyferencję obiektów. Przykłady: euklidesowa, Manhattan, kosinusowa.
23. Graf, ścieżka, cykl i droga w grafie
Graf = zbiór wierzchołków i krawędzi. Ścieżka = ciąg wierzchołków połączonych krawędziami. Cykl = ścieżka zamknięta. Droga = ścieżka bez powtarzających się wierzchołków.
24. Klasa równoważności
To zbiór elementów, które względem pewnej relacji są sobie równoważne. Przykład: podział liczb całkowitych według reszty z dzielenia.
25. Problemy programowania równoległego
Interferencja wątków, brak synchronizacji, zakleszczenia, głodzenie procesów.
26. Switch a router – różnice
Switch działa w warstwie 2 (łącza danych), przesyła pakiety w obrębie jednej sieci LAN. Router działa w warstwie 3 (sieć), łączy różne sieci i kieruje ruchem.
28. Co to jest grupowanie danych i podać algorytmy grupowania danych
Polega na przypisaniu obiektów do z góry zdefiniowanych klas. Algorytmy: drzewa decyzyjne, SVM, k-NN, regresja logistyczna.
29. Jak mierzyć zależność między atrybutami
Korelacja Pearsona, współczynnik Spearmana, współczynnik determinacji R², miary informacji wzajemnej.
30. Proces normalizacji modelu relacyjnego
Polega na przekształcaniu relacji, by eliminować redundancję i anomalie aktualizacji. Obejmuje kolejne postaci normalne (1NF, 2NF, 3NF, BCNF).
31. Maszyna Turinga
Abstrakcyjny model obliczeń z taśmą nieskończoną i głowicą czytająco-piszącą. Udowadnia, co jest obliczalne.
33. Numeryczne metody rozwiązywania układów równań
Metoda Gaussa, iteracyjna Jacobiego, Gauss-Seidla. Przybliżają rozwiązanie, gdy analityczne jest trudne.
34. Model relacyjny bazy danych
Opiera się na tabelach (relacjach) z wierszami (krotkami) i kolumnami (atrybutami). Klucze określają jednoznaczność i powiązania.
36. Wpływ indeksów na wydajność bazy danych
Przyspieszają wyszukiwanie danych, ale spowalniają operacje modyfikacji (INSERT/UPDATE/DELETE).
37. Klasyfikacja indeksów w bazie danych
Indeksy: podstawowe, złożone, unikalne, klastrowe i nieklastrowe. Przykład: indeks B-drzewa umożliwia szybkie wyszukiwanie zakresowe.
38. Metoda ścieżki krytycznej (CPM)
Służy do planowania projektów – identyfikuje najdłuższą ścieżkę zadań i czas minimalny realizacji projektu.
39. Metoda PERT
Stosuje analizę statystyczną czasu trwania zadań (optymistyczny, pesymistyczny, najbardziej prawdopodobny). Używana w projektach o dużej niepewności.
40. Metoda EVM (Earned Value Management)
Technika kontroli projektów. Łączy koszt, czas i zakres, mierząc wykonanie względem planu (SPI, CPI).
42. Układy liniowe i nieliniowe
Liniowe – spełniają zasadę superpozycji. Nieliniowe – występują nieliniowe zależności między zmiennymi.
43. Numeryczne metody całkowania
Metoda trapezów, prostokątów, Simpsona. Przybliżają całkę oznaczoną przez sumowanie pól figur.
44. Zasada działania maszyny Turinga
Na podstawie stanu i symbolu z taśmy wykonuje akcję: zapis, przesunięcie głowicy, zmiana stanu. Może symulować każdy algorytm.
45. Cechy systemów czasu rzeczywistego
Reakcja na zdarzenia w określonym czasie. Twarde RT wymagają gwarancji czasowych, miękkie RT tolerują opóźnienia.
46. Pojęcia w BPMN (bramka, zdarzenie, tor, basen)
Bramka – punkt decyzyjny (np. XOR, AND).
Zdarzenie – początek, koniec lub coś pośredniego w procesie.
Tor (lane) – rola/uczestnik procesu.
Basen (pool) – cała organizacja/proces główny.
47. Przeciążanie metod i operatorów
To definiowanie kilku metod o tej samej nazwie, ale różnych parametrach. W niektórych językach (C++, Python) można też przeciążać operatory (np. + dla własnej klasy).
48. Typ w językach programowania
Określa rodzaj przechowywanych danych i dopuszczalne operacje. Dzielimy na proste (int, bool) i złożone (tablice, obiekty).
49. Wielodziedziczenie w Javie
Java nie wspiera wielodziedziczenia klas, ale można osiągnąć podobny efekt przez interfejsy i klasy abstrakcyjne.
50. Diagramy UML
Strukturalne: diagram klas, komponentów, obiektów.
Behawioralne: przypadków użycia, sekwencji, aktywności, stanów.
51. Model sieciowy OSI
7 warstw: fizyczna, łącza danych, sieciowa, transportowa, sesji, prezentacji, aplikacji. Ułatwia standaryzację protokołów.


🔹 Model OSI (Open Systems Interconnection)
To teoretyczny model odniesienia opracowany przez ISO, żeby opisać, jak systemy komputerowe komunikują się w sieci.
Dzieli komunikację na 7 warstw, każda realizuje określone zadania:
Fizyczna – sygnały elektryczne, fale, kable, złącza (Ethernet, Wi-Fi na poziomie sygnału).
Łącza danych (Data Link) – przesył ramek między urządzeniami w tej samej sieci. Adresacja MAC, wykrywanie błędów (Ethernet, Wi-Fi, PPP).
Sieciowa (Network) – adresowanie logiczne i routowanie między sieciami (IP, ICMP).
Transportowa (Transport) – niezawodność, podział na segmenty, kontrola błędów, kolejność danych (TCP, UDP).
Sesji (Session) – zarządzanie sesjami (nawiązywanie, utrzymywanie, kończenie).
Prezentacji (Presentation) – tłumaczenie danych (szyfrowanie, kompresja, konwersja formatów, np. SSL/TLS).
Aplikacji (Application) – interfejs dla użytkownika/aplikacji (HTTP, FTP, SMTP, DNS).
👉 Model OSI jest abstrakcyjny – bardziej służy do nauki i standaryzacji niż do praktycznej implementacji.
🔹 Model TCP/IP
To praktyczna implementacja używana w internecie.
Ma mniej warstw (4), ale pokrywa funkcjonalnie te z OSI:
Dostępu do sieci (Network Access) – odpowiada warstwom 1–2 OSI (fizyczna + łącza danych).
Internet – odpowiada warstwie 3 OSI (protokoły: IP, ICMP).
Transport – odpowiada warstwie 4 OSI (TCP, UDP).
Aplikacji – odpowiada warstwom 5–7 OSI (HTTP, FTP, SMTP, DNS).
🔹 OSI vs TCP/IP – porównanie
OSI	TCP/IP	Przykłady protokołów
7. Aplikacji	Aplikacji	HTTP, FTP, SMTP, DNS
6. Prezentacji	Aplikacji	SSL/TLS, JPEG, ASCII
5. Sesji	Aplikacji	RPC, NetBIOS
4. Transportowa	Transport	TCP, UDP
3. Sieciowa	Internet	IP, ICMP
2. Łącza danych	Dostępu do sieci	Ethernet, Wi-Fi
1. Fizyczna	Dostępu do sieci	Kabel, sygnały

52. Protokoły połączeniowe i bezpołączeniowe
Połączeniowy – zestawia sesję (np. TCP).
Bezpołączeniowy – wysyła pakiety niezależnie (np. UDP).
53. Protokoły TCP, IP, ARP
TCP – niezawodna komunikacja strumieniowa.
IP – adresacja i routowanie pakietów.
ARP – zamiana adresów IP na MAC.
54. PKI – infrastruktura klucza publicznego
System zarządzania kluczami publicznymi i certyfikatami. Gwarantuje uwierzytelnianie, szyfrowanie i integralność.
55. Podpis elektroniczny
Szyfrowana wartość tworzona kluczem prywatnym. Odbiorca weryfikuje podpis za pomocą klucza publicznego.
56. Klasyfikacja metod testowania oprogramowania
Statyczne (przeglądy, analiza kodu).
Dynamiczne (testy jednostkowe, integracyjne, systemowe, akceptacyjne).
Manualne i automatyczne.
57. Wzorzec projektowy
Sprawdzone rozwiązanie typowego problemu projektowego. Przykład: Singleton (jeden obiekt w systemie), Observer (powiadamianie wielu obiektów o zmianie stanu).
58. Przestrzeń kolorów
Sposób reprezentacji barw. Przykłady: RGB (ekrany), CMYK (druk), HSV (intuicyjny opis koloru).
59. Dziedziczenie, generalizacja, enkapsulacja, asocjacje
Dziedziczenie – przejmowanie cech klasy bazowej.
Generalizacja – tworzenie klasy ogólniejszej.
Enkapsulacja – ukrywanie szczegółów implementacji.
Asocjacje – powiązania między klasami.
60. Miary w eksploracji danych (pozycyjne)
Środek rozkładu: mediana, kwartyle, decyle, percentyle. Służą do opisu danych i porównywania zbiorów.
61. Problemy mapowania danych obiektowych na relacyjne (niedopasowanie impedancji)
Obiektowe klasy i relacyjne tabele różnią się modelem danych (dziedziczenie, kolekcje, relacje wiele-do-wielu). Powoduje to trudności w ORM i spadek wydajności.
62. Struktury przechowywania danych
lista (dynamiczna kolekcja elementów),
stos (LIFO),
kolejka (FIFO),
tablica (indeksowany zbiór),
drzewo (hierarchia),
graf (sieć powiązań).



63. Drzewo decyzyjne
Model klasyfikacji/regresji, gdzie kolejne pytania dzielą dane na grupy. Używane w ML i systemach ekspertowych.
64. Algorytmy uczenia pod nadzorem – regresja
Algorytmy uczą się na danych z etykietami. Regresja przewiduje wartość liczbową (np. cena mieszkania), np. regresja liniowa.
66. Napięcie skuteczne
To wartość napięcia prądu zmiennego odpowiadająca napięciu stałemu dającemu tę samą moc cieplną. W sieci domowej 230 V.
67. Prąd
Uporządkowany ruch ładunków elektrycznych. Jednostka: amper.
68. Architektura von Neumanna
Komputer ma pamięć wspólną dla danych i programów, jednostkę sterującą, arytmetyczno-logiczną i urządzenia wejścia/wyjścia.
69. Klasa równoważności (w testowaniu)
Zbiór danych wejściowych traktowanych jako równoważne do testu. Redukuje liczbę przypadków testowych.
70. Twierdzenie Shannona o próbkowaniu
Aby wiernie odtworzyć sygnał, trzeba próbkować go z częstotliwością co najmniej 2× większą niż jego pasmo (fs ≥ 2fmax).
71. Automat skończony
Model obliczeń złożony ze stanów i przejść. Używany np. do analizy języków formalnych.
72. Zadanie optymalizacji
Polega na znalezieniu wartości zmiennych minimalizujących lub maksymalizujących daną funkcję celu przy określonych ograniczeniach.
73. Optymalizacja wklęsła i wypukła
Wypukła – każde minimum lokalne jest globalne (łatwiejsza).
Wklęsła – odpowiednik dla maksymalizacji.
74. Język bezkontekstowy
Język, którego gramatykę można opisać produkcjami typu A → α, gdzie A to nieterminal, α to dowolny ciąg symboli. Przykład: składnia wielu języków programowania.
75. Gramatyka, składnia, alfabet języka
Alfabet – zbiór symboli.
Składnia – reguły budowania poprawnych wyrażeń.
Gramatyka – formalny opis składni (np. gramatyka bezkontekstowa).
76. Rodzaje gramatyk – kontekstowe i bezkontekstowe
Bezkontekstowe – reguły zależą tylko od jednego nieterminala (np. języki programowania).
Kontekstowe – reguły mogą zależeć od otoczenia symbolu (np. niektóre języki naturalne).
77. Kompilacja i budowa kompilatora
Kompilacja = tłumaczenie kodu źródłowego na kod maszynowy. Etapy: analiza leksykalna, składniowa, semantyczna, optymalizacja, generacja kodu.
78. Informacja i entropia
Informacja to redukcja niepewności. Entropia Shannona mierzy średnią ilość informacji w zdarzeniu losowym.
79. Procesor a mikrokontroler
Procesor (CPU) – jednostka obliczeniowa komputera.
Mikrokontroler – CPU + pamięć + peryferia w jednym układzie (używany w urządzeniach wbudowanych).
80. System czasu rzeczywistego
Odpowiada na zdarzenia w określonym czasie. Twarde RT – opóźnienia niedopuszczalne, miękkie RT – możliwe niewielkie opóźnienia.
81. Notacja dużego O
Oznacza asymptotyczną złożoność algorytmu (górne ograniczenie). „O” pochodzi od niem. „Ordnung” = rząd wielkości.
82. Rodzaje złożoności algorytmu
Czasowa (czas wykonania).
Pamięciowa (zużycie pamięci).
Komunikacyjna (dla algorytmów rozproszonych).
83. Napięcie w sieci domowej
Skuteczne: 230 V.
Szczytowe: ok. 325 V.
Średnie: 0 (dla AC symetrycznego).
84. Jak odwrócić macierz
Stosując macierz dopełnień algebraicznych, eliminację Gaussa lub rozkład LU.
85. Programowanie strukturalne vs obiektowe
Strukturalne – procedury i funkcje, dane oddzielone od kodu.
Obiektowe – dane + metody w klasach, podejście bardziej zbliżone do rzeczywistości.
86. Algorytmy przeszukiwania w grafie
BFS (w szerz) – używa kolejki.
DFS (w głąb) – używa stosu/rekurencji.
87. Własności grafu
Skierowany/nieskierowany, ważony/nieważony, spójny, cykliczny/acykliczny, planarny.
88. Obserwator (automatyka)
Układ oceniający stan wewnętrzny systemu na podstawie obserwowanych sygnałów wyjściowych.
89. Kinematyka prosta i odwrotna (robotyka)
Prosta – obliczanie pozycji końcówki robota z wartości przegubów.
Odwrotna – obliczanie przegubów potrzebnych do osiągnięcia zadanej pozycji.
90. Warunki stabilności układu liniowego
Układ jest stabilny, gdy jego odpowiedź na ograniczone wejście jest ograniczona (BIBO stability). Matematycznie – wszystkie pierwiastki transmitancji muszą leżeć w lewej półpłaszczyźnie (ciągły) lub wewnątrz okręgu jednostkowego (dyskretny).
106. Klucz główny w modelu relacyjnym
Atrybut (lub ich zestaw), który jednoznacznie identyfikuje każdy wiersz w tabeli. Musi być unikalny i niepusty.
107. Typy bramek w BPMN
XOR (wykluczająca), OR (alternatywa), AND (równoległa), Event-based (na zdarzenia), Complex (złożona).
108. Prawa Kirchhoffa (prądowe i napięciowe)
Prądowe: suma prądów wpływających do węzła = suma wypływających.
Napięciowe: suma napięć w zamkniętej pętli = 0.
109. Zasada pracy maszyny Turinga
Na podstawie stanu i symbolu decyduje: zapisz/usuń symbol, przesuń głowicę, zmień stan. Dzięki temu symuluje każdy algorytm.
110. Operator liniowy
Funkcja między przestrzeniami wektorowymi, zachowująca dodawanie i mnożenie przez skalar.
111. Elementarne operacje w algorytmie genetycznym
Selekcja, krzyżowanie (crossover), mutacja, ewaluacja osobników.
112. Stos i jego rola w programowaniu
Struktura LIFO. Wykorzystany do wywołań funkcji, zapamiętywania powrotów, obliczeń rekurencyjnych.
113. Algorytmy sortowania (przykład)
Bąbelkowe (proste, O(n²)),
Quicksort (dziel i zwyciężaj, O(n log n)),
Merge sort (scalanie, O(n log n)).
114. Algorytm k-średnich (dr Twardy)
Wybiera k punktów początkowych, przypisuje obiekty do najbliższego centrum, iteracyjnie aktualizuje centra aż do zbieżności.
115. Problemy wielowątkowości (prof. Sierociuk)
Interferencja danych, synchronizacja, zakleszczenia, głodzenie procesów.
116. Rekurencja (WD)
Funkcja wywołująca samą siebie. Używana do problemów dzielonych na podproblemy (np. obliczanie Fibonacciego).
117. Indeks w bazie danych
Struktura przyspieszająca wyszukiwanie (np. B-drzewo, hash). Zwiększa szybkość SELECT, spowalnia UPDATE/INSERT.
118. Interpolacja i ekstrapolacja (prof. Malesza)
Interpolacja – przewidywanie wartości między znanymi punktami.
Ekstrapolacja – przewidywanie poza zakresem danych.
119. Diagramy UML w modelowaniu procesów biznesowych (prof. Dzieliński)
Diagram aktywności, przypadków użycia, sekwencji, klas – stosowane do opisu przepływu procesów.
120. Interpolacja funkcji (dr Twardy)
Znajdowanie funkcji, która przyjmuje znane wartości w danych punktach i przybliża wartości pośrednie (np. interpolacja liniowa, Newtona).
121. Perceptron wielowarstwowy i uczenie sieci neuronowych
Sieć z warstwą wejściową, ukrytą i wyjściową. Uczy się metodą propagacji wstecznej błędu i optymalizacji wag (np. gradient prosty, Adam).
122. Różnica między kompresją stratną a bezstratną
Bezstratna – dane można odtworzyć dokładnie (ZIP, PNG).
Stratna – część informacji tracona dla redukcji rozmiaru (MP3, JPEG).
123. Wzorzec projektowy i przykład architektonicznego
Wzorzec – powtarzalne rozwiązanie problemu projektowego. Architektoniczny przykład: MVC (Model-View-Controller).
124. Problemy programowania wielowątkowego (Dominik Olszewski)
Interferencja wątków.
Synchronizacja dostępu.
Głodzenie i zakleszczenia.
125. Role i zadania w SCRUM (KM)
Product Owner – zarządza backlogiem.
Scrum Master – wspiera zespół.
Developers – realizują zadania.
126. Indeks w bazie danych – pojęcie i klasyfikacja (WD)
Indeks to struktura przyspieszająca wyszukiwanie. Rodzaje: klastrowe, nieklastrowe, unikalne, złożone, pełnotekstowe.
127. Dziedziczenie w programowaniu obiektowym (DS)
Pozwala tworzyć nowe klasy na bazie istniejących. Sprzyja ponownemu wykorzystaniu kodu i hierarchii klas.
128. Perceptron – co to jest i do czego służy (MI)
Najprostsza sieć neuronowa – liniowy klasyfikator. Używany do zadań separowalnych liniowo.
129. Metody pomiaru projektu (EVM)
EVM porównuje planowany postęp z rzeczywistym. Wskaźniki: SPI (harmonogram), CPI (koszt).
130. Działy uczenia maszynowego (MI)
Uczenie nadzorowane (regresja, klasyfikacja).
Nienadzorowane (klasteryzacja).
Wzmacniające (reinforcement learning).
131. Maszyna Turinga (JSz)
Abstrakcyjny model obliczeń – taśma, głowica i zestaw reguł. Definiuje granice obliczalności.
132. Zarządzanie ryzykiem w projektach (WD)
Identyfikacja, analiza, planowanie reakcji i monitorowanie ryzyk. Strategie: unikanie, minimalizacja, transfer, akceptacja.
133. Graf (WM)
Zbiór wierzchołków połączonych krawędziami. Używany w informatyce m.in. w wyszukiwaniu tras, analizie sieci.
134. Automat skończony (JSz)
Model obliczeń z ograniczoną liczbą stanów i przejść. Wykorzystany np. w kompilatorach.
135. Metody przekazywania parametrów do funkcji (MM)
Przez wartość (kopiowanie).
Przez referencję/adres (zmiana oryginału).
Mieszane (np. const reference w C++).
136. Interpolacja i ekstrapolacja (WM)
Interpolacja – szacowanie wartości pomiędzy znanymi punktami.
Ekstrapolacja – przewidywanie wartości poza zakresem danych.
137. Transakcja w bazie danych (WD)
To zbiór operacji traktowany jako jedna całość. Musi spełniać warunki ACID (atomowość, spójność, izolacja, trwałość).
138. Zadanie programowania liniowego (AD)
Optymalizacja funkcji liniowej (np. zysk, koszt) przy liniowych ograniczeniach. Rozwiązywane np. metodą simpleks.
139. Encja i jej realizacja w bazie danych (MM)
Encja = obiekt rzeczywisty/pojęciowy. W bazie relacyjnej reprezentowana jako tabela.
140. Perceptron wielowarstwowy i algorytmy uczenia (AD)
Sieć z warstwami ukrytymi. Uczenie: propagacja wsteczna błędu, gradient prosty, optymalizatory (Adam, SGD).
141. Wzorce projektowe
Opisują sprawdzone rozwiązania problemów projektowych. Przykłady: Singleton, Observer, Factory, MVC.
142. Normalizacja modelu relacyjnego (WD)
Proces eliminacji redundancji i anomalii w danych przez kolejne postaci normalne (1NF, 2NF, 3NF, BCNF).
143. Algorytm genetyczny (KH)
Heurystyka inspirowana ewolucją. Operuje na populacji rozwiązań, stosuje selekcję, krzyżowanie, mutacje.
144. Paradygmat programowania obiektowego – kluczowe pojęcia (RŁ)
Klasy, obiekty, dziedziczenie, polimorfizm, enkapsulacja, abstrakcja.
145. Wzorzec projektowy i przykład architektoniczny (WD)
Wzorzec = powtarzalne rozwiązanie problemu. Architektoniczny przykład: warstwowa architektura systemu.
146. Złożoność obliczeniowa (KH)
Miara zasobów (czas, pamięć) potrzebnych do wykonania algorytmu względem rozmiaru danych.
147. Regresja vs klasyfikacja (RŁ)
Regresja – przewiduje wartości liczbowe (np. ceny).
Klasyfikacja – przypisuje obiekty do kategorii (np. spam/nie-spam).
148. Mechanizm transakcji i postulaty ACID (WD)
Atomowość – całość lub nic.
Consistency – dane spójne.
Isolation – transakcje niezależne.
Durability – trwałość po zapisie.
149. Interpolacja (WM)
Przybliżanie wartości funkcji między znanymi punktami. Np. interpolacja liniowa, wielomianowa.
150. Jakość i metryka jakości w zarządzaniu projektem IT (KH)
Jakość = zgodność z wymaganiami. Metryki: liczba błędów, pokrycie testami, wskaźniki wydajności (CPI, SPI).
151. Ekstrapolacja (WM)
Szacowanie wartości funkcji poza zakresem znanych danych. Bardziej ryzykowna niż interpolacja, bo opiera się na ekstrapolacji trendu.
152. Sieć neuronowa (SS)
Model matematyczny inspirowany mózgiem. Składa się z neuronów i warstw. Używana do klasyfikacji, rozpoznawania obrazów, NLP.
153. Ścieżka krytyczna w projektach (WD)
Najdłuższa sekwencja zadań, która wyznacza minimalny czas realizacji projektu. Opóźnienie któregokolwiek zadania opóźnia cały projekt.
154. Grafy do ścieżki krytycznej (WD)
Wykorzystuje się grafy acykliczne skierowane (DAG), gdzie wierzchołki to zadania, a krawędzie to zależności.
155. Narzędzia modelowania procesów biznesowych (AD)
BPMN, UML, EPC, IDEF0, ARIS, Bizagi, Camunda.
156. Semantyka języka BPMN (WD)
Opis znaczenia elementów BPMN: zdarzeń, bramek, basenów, torów, aktywności. Umożliwia jednoznaczne zrozumienie procesów.
157. Maszyna Turinga (DS)
Abstrakcyjny model obliczeń (taśma, głowica, stany). Podstawa teorii obliczalności.
158. Transakcja i warunki (MT)
Transakcja to jednostka pracy w bazie. Warunki ACID: atomowość, spójność, izolacja, trwałość.
159. Atak DoS (DS)
Denial of Service – przeciążenie systemu nadmiarem żądań, co uniemożliwia obsługę prawidłowych użytkowników.
160. Metoda EVM w projektach (WD)
Porównuje planowany a rzeczywisty postęp. Wskaźniki: SPI (czas), CPI (koszt).
161. Cztery paradygmaty programowania obiektowego (WC)
Abstrakcja, dziedziczenie, polimorfizm, enkapsulacja.
162. Pamięć cache (JS)
Szybka pamięć pośrednia między RAM a CPU. Przyspiesza dostęp do często używanych danych.
163. Gramatyka bezkontekstowa (JS)
Reguły w postaci A → α (nieterminal → ciąg symboli). Stosowana w opisach języków programowania.
164. Problem komiwojażera (JS)
Znalezienie najkrótszej drogi odwiedzającej wszystkie miasta i wracającej do startu. Klasyczny problem NP-trudny.
165. Numeryczne wyliczanie całki oznaczonej (DS)
Metody: prostokątów, trapezów, Simpsona. Przybliżają pole pod krzywą.
166. Graf (WM)
Struktura danych: zbiór wierzchołków połączonych krawędziami. Wykorzystywana w sieciach, grafach społecznych, optymalizacji tras.
167. Cechy transakcji w bazie (DS)
Transakcja musi spełniać zasady ACID: atomowość, spójność, izolacja, trwałość.
168. Zadanie optymalizacji (MT)
Polega na znalezieniu najlepszego rozwiązania (minimum/maksimum) dla funkcji celu przy zadanych ograniczeniach.
169. Przestrzenie barw (GS)
Sposób reprezentacji kolorów: RGB (monitory), CMYK (druk), HSV (intuicyjne kolory), CIE Lab (percepcyjna różnica barw).
170. Cechy transakcji (MT)
Transakcja musi być atomowa, spójna, izolowana i trwała – gwarantuje poprawność danych.
171. Regresja vs klasyfikacja (RŁ)
Regresja – przewiduje wartości liczbowe.
Klasyfikacja – przypisuje obiekty do kategorii.
172. Kodowanie stratne i bezstratne (MI)
Bezstratne – ZIP, PNG.
Stratne – JPEG, MP3. Różnią się możliwością odtworzenia danych.
173. Bezpieczeństwo w sieci komputerowej (MI)
Stosowanie zapór (firewall), szyfrowania, VPN, systemów IDS/IPS, certyfikatów i haseł.
174. Przeuczenie sieci neuronowej (DO)
Model nauczył się „na pamięć” danych treningowych i nie generalizuje. Rozwiązania: regularyzacja, dropout, więcej danych.
175. Paradygmat obiektowego programowania (RŁ)
Opiera się na klasach i obiektach. Kluczowe cechy: enkapsulacja, dziedziczenie, polimorfizm, abstrakcja.
176. Protokół dwufazowy (2PC) w bazach (WD)
Koordynator pyta uczestników o gotowość (faza 1). Jeśli wszyscy potwierdzą – zatwierdza (faza 2). Zapewnia spójność w systemach rozproszonych.
177. Obiekt pozorny (mock) w testowaniu (DO)
Sztuczny obiekt symulujący działanie zależności (np. bazy danych). Ułatwia testy jednostkowe.
178. Klasyfikacja wymagań na system (MS)
Funkcjonalne (co system ma robić).
Niefunkcjonalne (wydajność, bezpieczeństwo).
Biznesowe, użytkownika, systemowe.
179. Prąd (MW)
Uporządkowany ruch ładunków elektrycznych. Jednostka: amper (A).
180. Interpretacja pochodnej i całki oznaczonej (MW)
Pochodna – szybkość zmian (np. prędkość to pochodna drogi).
Całka – pole pod wykresem (np. droga jako całka z prędkości).
181. Normalizacja w bazach danych – kiedy stosować / nie stosować (KH)
Stosujemy: aby wyeliminować redundancję i anomalie (wstawiania, usuwania, aktualizacji).
Nie stosujemy (lub częściowo): w hurtowniach danych, gdzie liczy się szybkość odczytu, a nie spójność.
182. Encja i jej zastosowanie
Encja = obiekt rzeczywisty/pojęciowy. W bazach danych odpowiada tabeli, a jej wystąpienia to wiersze.
183. Stos w programowaniu (WD)
Struktura LIFO. Używana do obsługi wywołań funkcji, rekurencji i zapamiętywania stanów.
184. Język BPMN i jego konstrukcje (WD)
BPMN = standard notacji procesów biznesowych. Konstrukcje: zdarzenia, bramki, aktywności, baseny, tory.
185. Złożoność obliczeniowa (KMk)
Miara zasobów (czas, pamięć) potrzebnych do wykonania algorytmu w zależności od danych wejściowych.
186. Graf i zastosowania w informatyce (WD)
Grafy stosuje się do analizy sieci komputerowych, tras w GPS, grafów społecznościowych, kompilatorów.
187. Model relacyjny baz danych (mm)
Opiera się na tabelach (relacjach). Wiersze = krotki, kolumny = atrybuty. Klucze definiują powiązania.
188. Złożoność obliczeniowa (kmk)
To samo co wyżej – mierzona czasowo, pamięciowo, komunikacyjnie.
189. Cechy systemów rozproszonych (gS)
Przezroczystość (lokalizacji, dostępu, replikacji), skalowalność, odporność na awarie, współbieżność.
190. OLTP vs OLAP (kmk)
OLTP – systemy transakcyjne (operacje bieżące).
OLAP – analityczne, wspierające podejmowanie decyzji (hurtownie danych).
191. Język programowania i klasyfikacja (WD)
Formalny system zapisu algorytmów. Klasyfikacja: imperatywne, deklaratywne, proceduralne, obiektowe, funkcyjne, logiczne.
192. Różnica switch vs router (gS)
Switch – działa w warstwie 2, łączy urządzenia w sieci lokalnej.
Router – działa w warstwie 3, łączy różne sieci i kieruje ruchem.
193. Model OSI (KMk)
7 warstw: fizyczna, łącza danych, sieciowa, transportowa, sesji, prezentacji, aplikacji.
194. Role i wydarzenia w Scrum (KMk)
Role: Product Owner, Scrum Master, Zespół Developerski.
Wydarzenia: Sprint, Daily, Sprint Planning, Review, Retrospektywa.
195. Urządzenie cyfrowe vs analogowe (MS)
Cyfrowe – przetwarza sygnały w postaci 0 i 1 (komputery).
Analogowe – operuje na sygnałach ciągłych (radio, wzmacniacz).
196. Konstruktor i jego rola
Specjalna metoda klasy wywoływana przy tworzeniu obiektu. Służy do inicjalizacji pól i przygotowania obiektu do użycia.
197. Obiekt pozorny w testowaniu (DO)
Sztuczny obiekt (mock/stub/fake) zastępujący prawdziwy komponent w testach. Umożliwia testowanie w izolacji.
198. Język programowania – pojęcie i klasyfikacja (WD)
Język do zapisu algorytmów. Klasyfikacja: imperatywne, deklaratywne, obiektowe, proceduralne, funkcyjne, logiczne, skryptowe.
199. Różnica BPMN vs UML (KMk)
BPMN – do modelowania procesów biznesowych.
UML – do modelowania systemów informatycznych.
UML bardziej techniczny, BPMN biznesowy.
200. Diagram klas w UML (DO)
Przedstawia klasy, ich atrybuty, metody i relacje (dziedziczenie, asocjacje, kompozycje).
201. Ścieżka krytyczna w projektach (WD)
Najdłuższy ciąg zadań zależnych czasowo. Wyznacza minimalny czas realizacji projektu.
202. Role i wydarzenia w Scrum (KMk)
Role: Product Owner, Scrum Master, Developers.
Wydarzenia: Sprint, Daily, Planning, Review, Retrospektywa.
203. Sieci neuronowe z pamięcią (DO)
Sieci rekurencyjne (RNN, LSTM, GRU) – mają zdolność zapamiętywania sekwencji i historii.
204. Inne modele sieci (niż transformery)
Sieci bayesowskie,
Bag of Words,
Modele tematyczne (LDA),
Starsze sieci CNN i RNN.
205. Metoda EVM (WD)
Earned Value Management – łączy koszty, czas i zakres. Wskaźniki: SPI (czas), CPI (koszt).
206. Transakcja w bazach danych (MT)
Jednostka operacji na danych spełniająca zasady ACID. Chroni spójność i integralność danych.
207. Normalizacja w bazie danych (KMk)
Proces rozbijania tabel, aby eliminować redundancję i anomalie. Stosuje kolejne postaci normalne (1NF, 2NF, 3NF, BCNF).
208. Algorytmy sortowania (WD)
Bubble sort, Quick sort, Merge sort, Heap sort. Różnią się złożonością i podejściem.
209. Transakcja i warunki (MT)
Musi być atomowa, spójna, izolowana i trwała – ACID.
210. Warunek odtworzenia sygnału cyfrowego (WD)
Twierdzenie Shannona: częstotliwość próbkowania musi być ≥ 2 × najwyższa częstotliwość sygnału.
211. Model OSI (KMk)
Standard komunikacji sieciowej w 7 warstwach: fizyczna, łącza danych, sieciowa, transportowa, sesji, prezentacji, aplikacji.
212. Transakcja (MT)
Spójny zbiór operacji w bazie, który musi spełniać ACID.
213. Topologie sieci neuronowych (DO)
Jednowarstwowe,
Wielowarstwowe (MLP),
Rekurencyjne (RNN, LSTM),
Konwolucyjne (CNN).
214. Maszyna Turinga (MT)
Abstrakcyjny model obliczeń – taśma, głowica, zestaw stanów i reguł. Fundament teorii obliczalności.
215. Transakcja i różnica od procedury (Maciej Sławiński)
Transakcja = jednostka pracy w bazie (ACID). Procedura = zestaw instrukcji SQL, który może obejmować transakcje.
216. Metody modelowania procesów biznesowych (AD)
BPMN, UML, EPC, IDEF0, ARIS.
217. Metody badania złożoności algorytmów (WD)
Analiza asymptotyczna – notacja O, Ω, Θ. Duże O = górne ograniczenie, małe o = ścisłe ograniczenie.
218. Obiekty pozorne (Java, DO)
Mocki/stuby – używane w testach jednostkowych do symulowania zależności.
219. Paradygmat obiektowy (KH)
Opiera się na klasach i obiektach. Kluczowe cechy: abstrakcja, dziedziczenie, polimorfizm, enkapsulacja.
220. Sieci neuronowe – pojęcie i zastosowania (AD)
Modele AI wzorowane na mózgu. Zastosowania: NLP, rozpoznawanie obrazów, prognozy, gry.
221. Postaci normalne relacji (KH)
1NF – atomowość,
2NF – brak częściowej zależności od klucza,
3NF – brak zależności przechodnich.
223. Klasyfikacja języków programowania
Imperatywne, deklaratywne, proceduralne, obiektowe, funkcyjne, logiczne, skryptowe.
224. Złożoność obliczeniowa algorytmu (RŁO)
Miara czasu/pamięci potrzebnych do wykonania algorytmu.
225. Numeryczna metoda wyznaczania całki oznaczonej
Metody trapezów, prostokątów, Simpsona.
226. Inteligentne techniki obliczeniowe
Metody inspirowane naturą, np. sieci neuronowe, algorytmy genetyczne, logika rozmyta. Stosowane w AI i optymalizacji.
227. Warunki dla 3NF
Relacja w 2NF i brak zależności przechodnich między atrybutami niekluczowymi.
228. Klasa
Szablon opisujący strukturę (pola) i zachowanie (metody) obiektów. Na jej podstawie tworzy się obiekty.
229. Transakcje w systemach współbieżnych
Problem anomalii (dirty read, phantom read). Rozwiązania: poziomy izolacji, protokół 2PC.
230. Dziedziczenie w programowaniu obiektowym
Tworzenie nowych klas na podstawie istniejących. Ułatwia ponowne wykorzystanie kodu.
231. Proces uczenia sieci neuronowej
Dopasowywanie wag między neuronami poprzez minimalizację błędu (np. backpropagation).
232. Algorytm tyrana (Konrad Markowski)
Do wyboru koordynatora wybierany jest proces z najwyższym numerem.
233. Mechanizm wskaźnika w programowaniu
Przechowuje adres zmiennej. Pozwala manipulować danymi pośrednio.
234. Problem NP-zupełny
Trudne problemy, dla których nie znamy algorytmu wielomianowego. Każdy NP-problem można do nich sprowadzić.
235. System operacyjny – budowa
Jądro, zarządzanie pamięcią, procesami, system plików, sterowniki.
236. Pojęcie odległości
Miara różnicy między obiektami. Przykłady: euklidesowa, Manhattan, kosinusowa.
237. Graf, ścieżka, cykl, droga
Graf – wierzchołki i krawędzie.
Ścieżka – kolejne wierzchołki połączone krawędziami.
Cykl – ścieżka zamknięta.
Droga – ścieżka bez powtórzeń wierzchołków.
238. Klasa równoważności
Zbiór elementów równych względem danej relacji.
239. Problemy programowania równoległego
Synchronizacja, interferencja, zakleszczenia, głodzenie.
240. Różnica switch vs router
Switch – warstwa 2 (łącza danych), łączy urządzenia w LAN.
Router – warstwa 3 (sieć), łączy różne sieci i kieruje ruchem.
