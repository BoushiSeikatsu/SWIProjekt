# Iterační vývojový proces

Tento dokument popisuje iterativní proces vývoje softwaru, který tým používá k efektivní správě úkolů, sledování sprintů, řešení problémů a správě verzí kódu s využitím nástrojů jako **Git**, **CI/CD pipeline** a **Trello**.

## 1. Sledování průběhu sprintů

### 1.1 Kanban Board
- Používáme [Trello](https://trello.com/b/nXTrOgyw/základní-principy-sw-inženýrství) pro vizualizaci stavu úkolů v průběhu sprintu.
- **Sloupce Kanban boardu**:
  - **To Do**: Úkoly připravené k zahájení.
  - **In Progress**: Úkoly, na kterých se aktuálně pracuje.
  - **Done**: Dokončené úkoly.
- Každý člen týmu je odpovědný za pravidelnou aktualizaci stavu svých úkolů.

### 1.2 Týdenní Stand-up Meetingy
- **Týdenní stand-up**: Krátká schůzka (max. 5 minut) na začátku každého týdne.
- Každý člen týmu:
  - Shrne, co udělal od poslední schůzky.
  - Popíše, na čem bude pracovat tento týden.
  - Identifikuje případné problémy nebo blokace.
- Cílem těchto setkání je rychlá synchronizace a včasná identifikace potenciálních překážek.

## 2. Zajištění splnění User Stories

### 2.1 Definition of Done (DoD)
- Každá **user story** musí mít jasně definovanou **Definition of Done** (DoD), která specifikuje, kdy je úkol považován za dokončený.
- Příklad DoD může zahrnovat:
  - Úspěšné otestování všech funkcí.
  - Ověření, že byly splněny veškeré požadavky zadané v user story.
  - Review kódu a schválení týmem.

### 2.2 Řešení problémů
- Pokud při vývoji vznikne technický problém, člen týmu může požádat o pomoc jiného člena týmu.
- Závažnější problémy lze eskalovat na stand-up meetingu nebo přehodnotit prioritizaci úkolů na základě vzniklých blokací.

## 3. Správa verzí a integrace

### 3.1 Používání Git
- Pro správu verzí používáme **Git** a kód je hostován na **GitHubu**.
- Změny v kódu jsou spravovány prostřednictvím následujícího verzovacího schématu:
  - **Patch verze** (zvýšení na 0.0.x) při každém pushi do hlavní větve `develop`.
  - **Minor verze** (zvýšení na 0.x.0) při každém pushi do hlavní větve `main`.
- Vývoj probíhá lokálně ve feature větvích, které jsou následně pushovány do GitHubu pomocí příkazu:
  - `git push -u origin branch-name`
- Před sloučením změn do `main` nebo `develop` větve je nutné:
  1. Vytvořit **Pull Request**.
  2. Projít procesem **Code Review**.

### 3.2 Git Flow
Používáme metodologii **Git Flow**, která zahrnuje následující větve:
- **main**: Stabilní větev připravená k nasazení do produkce.
- **develop**: Větev pro integraci nových funkcí a jejich testování.
- **feature/us-xxx**: Každá nová funkce je vyvíjena v samostatné feature větvi.
- **hotfix/us-xxx**: Větev pro rychlé opravy kritických chyb na produkci.

### 3.3 CI/CD Pipeline
- CI/CD pipeline je implementována pomocí **GitHub Actions**:
  - Každý push do větví `develop` nebo `main` automaticky spouští pipeline.
  - Pipeline zajišťuje:
    - Automatické testování.
    - Sestavení a nasazení aplikace.
