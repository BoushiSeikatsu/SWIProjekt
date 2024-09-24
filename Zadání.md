# Funkční požadavky

1. **Správa skladových zásob**  
   - Aplikace umožní přidávat, upravovat a odstraňovat záznamy o skladových položkách (název, popis, množství, cena, dodavatel, kategorie).

2. **Sledování stavu zásob**  
   - Uživatelé mohou v reálném čase sledovat množství dostupných zásob na skladě a zobrazovat historii pohybu zásob (příjem, výdej, převody mezi sklady).

3. **Upozornění na nízké zásoby**  
   - Aplikace automaticky upozorní uživatele (notifikace, e-mail) při dosažení předem stanovené minimální hladiny zásob.

4. **Správa objednávek**  
   - Uživatelé mohou vytvářet, spravovat a sledovat objednávky u dodavatelů, včetně generování objednávkových dokladů a sledování stavu objednávek (vytvořeno, odesláno, doručeno).

5. **Historie transakcí a reporty**  
   - Aplikace bude zaznamenávat historii všech transakcí (příjemky, výdejky, objednávky) a umožní generovat reporty (zásoby, objednávky, pohyby).

6. **Role a oprávnění uživatelů**  
   - Aplikace bude podporovat různé role (správce, uživatel skladu) s rozdílnými úrovněmi oprávnění, aby bylo možné řídit přístup k jednotlivým funkcím.

7. **Integrace s externími systémy**  
   - Aplikace bude umožňovat export/import dat (např. ve formátech CSV, Excel) nebo integraci s ERP systémy.

# Nefunkční požadavky

1. **Dostupnost a škálovatelnost**  
   - Aplikace musí být dostupná 99,9 % času a musí být schopná zvládnout nárůst objemu skladových položek i počtu uživatelů bez výrazného zpomalení.

2. **Výkon**  
   - Aplikace musí reagovat do 2 sekund při běžných operacích (vyhledávání zásob, zobrazení seznamů), i při velkém počtu záznamů (např. 10 000 položek).

3. **Bezpečnost**  
   - Data o zásobách musí být šifrována a aplikace musí podporovat dvoufaktorovou autentizaci pro přístup uživatelů.

4. **Zálohování dat**  
   - Aplikace musí automaticky zálohovat data alespoň jednou denně a umožnit snadnou obnovu při selhání systému.

5. **Podpora různých zařízení**  
   - Aplikace musí být responzivní a plně funkční na desktopových i mobilních zařízeních.

6. **Uživatelská přívětivost (UX/UI)**  
   - Aplikace musí být intuitivní, s jasnou navigací a snadným ovládáním i pro méně zkušené uživatele.

7. **Rychlost nasazení změn**  
   - Aktualizace a nasazení nových verzí musí probíhat bez dlouhých odstávek (ideálně s nulovým výpadkem).
