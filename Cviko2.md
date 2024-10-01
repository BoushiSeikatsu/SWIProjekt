### Úkol: Vytvořte procesní postupy pro každou z následujících fází Agile vývoje na základě zadání od zákazníka:

#### 1. Sběr a správa požadavků (Product Backlog):
- Jak budete spravovat backlog požadavků, které jste dostali od zákazníka?
    - Určit které požadavky budou MVP a které mají největší prioritu. Takže rozdělení priorit. Budeme používat Giru.
- Jak budete komunikovat se zákazníkem o prioritách jednotlivých user stories?
    - Zeptáme se jich co bude nejdůležitější a co může počkat. 
- Jak často budete dělat revize backlogu a upravovat ho podle zpětné vazby?
    - Na konci sprintu
- Jak budete definovat Definition of Done pro každý inkrement?
    - Budeme mít cíl a pokud bude splněn.
    - Cíl je done, testy jsou napsané, dokumentace je taky napsaná 

#### 2. Plánování sprintů a Product Increments:
- Jak budete plánovat vývojové sprinty? Jakým způsobem budete rozdělovat práci mezi členy týmu?
    - Na konci předešlého sprintu. Tasky budeme mít rozdělené na oblasti, třeba (backend, frontend) a lidi si budou brát tasky sami.
- Jak budete určovat, které user stories budou součástí každého sprintu? Jaké techniky odhadování budete používat (např. story points)?
    - User stories vybírat závisle na tom co zrovna děláme.
    - User stories -> jsou to požadavky řečeny v kontextu. Z toho se pak odvodí tasky. 
    - To budeme modelovat v tom sprint planingu. 12 story pointů na sprint, počítali jsme to v porovnání s kredity. Jeden story point je půl dne. 
- Jak budete zajišťovat, že výstupy sprintu budou plně funkční a připravené na integraci?
    - Budeme mít acceptance test s default use casem.
- Jak budete organizovat sprint review a retrospektivy?
    - Na cviku každý řekne co chtěl stihnout a kolik toho reálně stihnul a koukneme se jestli jsme splnili definition of done. Podle toho zhodnotíme sprint.
    - Použít to jako poučení k dalšímu sprintu.

#### 3. Implementace (Iterace vývoje):
- Jak budete sledovat průběh jednotlivých sprintů? Jak budete řídit úkoly na denní bázi (daily stand-upy, Kanban)?
    - Budeme mít Kanban board, členové týmu se o to postarají. Weekly stand upy na začátku cvika (5 minut). 
- Jak budete zajišťovat, že user stories budou splněny a jak budete řešit problémy, které mohou vzniknout během sprintu?
    - Každé user story bude mít definition of done. Pokud to bude problém technického charakteru tak požádat o pomoc jiného člena týmu.
- Jak budete využívat nástroje pro správu verzí a integraci (např. Git, CI/CD pipeline)?
    - Git, každý push do main bude zvyšovat main verze, budeme mít rozdílné branche.
    - Můžeme implementovat Git Flow pro 4 branche, nebo budeme mít develop, main, feature a hotfix.

#### 4. Testování (Continuous Integration):
- Jak budete zajišťovat testování průběžně během vývoje? Jakým způsobem zajistíte, že každý inkrement je testovaný a otestovaný, než se nasadí?
    - Github testy, testy poběží při každém pushi, asi každý si bude psát svoje testy, budeme mít testy na code coverage 69%
- Jaké testy budete provádět během sprintu (unit, integrační, akceptační)?
    - Akceptační 
- Jak budete automatizovat testovací procesy a zajišťovat zpětnou vazbu (např. Jenkins, Selenium)?
    - Přes github

#### 5. Nasazení a retrospektiva:
- Jak budete plánovat nasazení po každém sprintu nebo inkrementu?
    - Sprint planing na konci sprintu, které user stories jsme nestihli a případně vymyslíme nějaké nové
- Jak budete monitorovat produkční prostředí sprintu? Jak budete řešit případné chyby? Jak budete řešit rollback v případě problémů?
    - Pokud naše aplikace by běžela například na serveru, tak například jestli běží server. Rollback -> V případě appky spustíme safety pipeline co se vrátí do starší verze.
    - Nebo to možná budeme řešit před Docker, pro databázi se vždy ve 3 ráno udělá kopie a nahraje se na cloud
- Jak budete organizovat retrospektivu po každém sprintu a jakým způsobem budete implementovat získanou zpětnou vazbu?
    - Organizace bude na cviku jednou za měsíc na začátku sprintu a implementovat to bude způsobem že na čem se shodneme
