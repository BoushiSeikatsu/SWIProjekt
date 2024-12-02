# Testování (Continuous Integration)

Testování je klíčovou součástí vývoje softwaru, která zajistí, že každá část kódu je dostatečně prověřená před nasazením. Continuous Integration (CI) umožňuje pravidelné testování v průběhu celého vývoje. V této části popíšeme, jakým způsobem budeme zajišťovat, aby byly všechny změny průběžně testovány a splňovaly požadovanou kvalitu.

## 1. Průběžné testování během vývoje

Průběžné testování zajišťuje, že jakmile vývojář provede jakoukoli změnu a odešle ji na vzdálený repozitář (např. na GitHubu), automaticky proběhne sada testů.
- **GitHub testy**: Testy budou spuštěny automaticky při každém `push`i kódu do repozitáře. To znamená, že jakákoliv změna bude okamžitě ověřena, aby nedošlo k narušení již existující funkcionality.
- **Odpovědnost za psaní testů**: Každý člen týmu bude odpovědný za psaní testů ke svému kódu, což zajistí, že každý nový kód bude přicházet spolu s odpovídajícími testovacími případy. Každý vývojář bude pokrývat své funkce nejen unit testy, ale také integračními testy tam, kde je to potřeba.
- **Code coverage**: Testování bude zacíleno na pokrytí kódu minimálně z 69 %. Code coverage je ukazatel, který měří, jak velká část kódu byla během testování skutečně vykonána. Toto číslo nám pomůže sledovat, zda testy pokrývají dostatečnou část projektu.

## 2. Druhy testů prováděné během sprintu

Během vývojového cyklu v rámci jednotlivých sprintů bude tým provádět různé typy testů:
- **Jednotkové testy (unit tests)**: Tyto testy se zaměřují na jednotlivé části kódu, jako jsou jednotlivé funkce nebo metody. Unit testy budou psány pro všechny nové části aplikace.
- **Integrační testy**: Tyto testy ověřují, zda různé moduly nebo části systému spolupracují správně. Integrační testy se budou používat při integraci nových funkcí do existujícího kódu.
- **Akceptační testy**: Na konci sprintu proběhnou akceptační testy, které zajišťují, že daná funkcionalita splňuje požadavky a očekávání uživatele. Tyto testy budou klíčové pro dokončení a uzavření jednotlivých user stories.

## 3. Automatizace testovacích procesů a zpětná vazba

Automatizace testování hraje důležitou roli v agilním vývoji, kde je potřeba neustále a rychle nasazovat nové verze softwaru. Automatické testy zajišťují, že se jakékoliv chyby objeví co nejdříve, což umožňuje rychlou reakci a opravu.
- **Automatizace přes GitHub Actions**: Testovací proces bude plně automatizován pomocí GitHub Actions, které zajistí, že po každém `push`i nebo `pull request`u proběhne sada testů. GitHub Actions budou nakonfigurovány tak, aby spouštěly všechny definované testy, včetně unit testů, integračních testů a testů na code coverage. Tento proces zajistí rychlou zpětnou vazbu a upozorní vývojáře, pokud některý test selže.
- **Monitoring výsledků testů**: Výsledky testů budou viditelné přímo v rámci GitHubu, což poskytne týmu okamžitý přehled o stavu kódu. V případě neúspěchu budou vývojáři ihned upozorněni, aby mohli rychle reagovat.

Tento přístup k testování a automatizaci v rámci CI zajišťuje, že kód je průběžně kontrolován a otestován. Tím minimalizujeme riziko chyb při nasazení a udržujeme vysokou kvalitu softwaru po celou dobu vývoje.
