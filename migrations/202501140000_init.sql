-- +goose Up
-- +goose StatementBegin
--trgm
CREATE SCHEMA trgm;
CREATE EXTENSION pg_trgm WITH SCHEMA trgm;
--groups
CREATE TABLE IF NOT EXISTS public.groups
(
    group_id uuid NOT NULL DEFAULT gen_random_uuid(),
    name character varying COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT groups_pk PRIMARY KEY (group_id)
)
TABLESPACE pg_default;
ALTER TABLE IF EXISTS public.groups
    OWNER to postgres;
CREATE INDEX IF NOT EXISTS groups_name_lower_trgm
    ON public.groups USING gist
    (lower(name::text) COLLATE pg_catalog."default" trgm.gist_trgm_ops)
    TABLESPACE pg_default;
CREATE UNIQUE INDEX IF NOT EXISTS groups_name_lower_unique_btree
    ON public.groups USING btree
    (lower(name::text) COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default;
--groups. insert demo data
INSERT INTO public.groups
VALUES 
('33ce21ac-1875-4420-9e95-e422922f5e6d','Пикник'),
('5176a448-10b9-465f-8e78-bb00502f9055','Nautilus Pompilius'),
('cac0f25f-e305-43d6-9155-e73f73b59817','Кино');
--songs
CREATE TABLE IF NOT EXISTS public.songs
(
    song_id uuid NOT NULL DEFAULT gen_random_uuid(),
    group_id uuid NOT NULL,
    name character varying COLLATE pg_catalog."default" NOT NULL,
    release_date date,
    text character varying COLLATE pg_catalog."default",
    link character varying COLLATE pg_catalog."default",
    CONSTRAINT songs_pk PRIMARY KEY (song_id),
    CONSTRAINT group_fk FOREIGN KEY (group_id)
        REFERENCES public.groups (group_id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE RESTRICT
        NOT VALID
)
TABLESPACE pg_default;
ALTER TABLE IF EXISTS public.songs
    OWNER to postgres;
CREATE INDEX IF NOT EXISTS songs_group_id_btree
    ON public.songs USING btree
    (group_id ASC NULLS LAST)
    TABLESPACE pg_default;
CREATE INDEX IF NOT EXISTS songs_link_lower_trgm
    ON public.songs USING gist
    (lower(link::text) COLLATE pg_catalog."default" trgm.gist_trgm_ops)
    TABLESPACE pg_default;
CREATE INDEX IF NOT EXISTS songs_name_lower_trgm
    ON public.songs USING gist
    (lower(name::text) COLLATE pg_catalog."default" trgm.gist_trgm_ops)
    TABLESPACE pg_default;
CREATE INDEX IF NOT EXISTS songs_release_date_btree
    ON public.songs USING btree
    (release_date ASC NULLS LAST)
    TABLESPACE pg_default;
CREATE INDEX IF NOT EXISTS songs_text_lower_trgm
    ON public.songs USING gist
    (lower(text::text) COLLATE pg_catalog."default" trgm.gist_trgm_ops)
    TABLESPACE pg_default;
--songs. insert demo data
INSERT INTO public.songs VALUES
('87b42eb1-5ba8-4662-96dc-0c4aa52e453c','cac0f25f-e305-43d6-9155-e73f73b59817','Звезда по имени Солнце','1989-01-01',E'Белый снег, серый лед,\nНа растрескавшейся земле.\nОдеялом лоскутным на ней -\nГород в дорожной петле.\n\nА над городом плывут облака,\nЗакрывая небесный свет.\nА над городом - желтый дым,\nГороду две тысячи лет,\nПрожитых под светом Звезды\nПо имени Солнце...\n\nИ две тысячи лет - война,\nВойна без особых причин.\nВойна - дело молодых,\nЛекарство против морщин.\n\nКрасная, красная кровь\nЧерез час уже просто земля,\nЧерез два на ней цветы и трава,\nЧерез три она снова жива\nИ согрета лучами Звезды\nПо имени Солнце...\n\nИ мы знаем, что так было всегда,\nЧто Судьбою больше любим,\nКто живет по законам другим\nИ кому умирать молодым.\n\nОн не помнит слово да и слово нет,\nОн не помнит ни чинов, ни имен.\nИ способен дотянуться до звезд,\nНе считая, что это сон,\nИ упасть, опаленным Зведой по имени Солнце','https://youtu.be/jQV5VXfKDYc'),
('f3e4d076-9ec0-4b90-a64f-8d08f6992e32','cac0f25f-e305-43d6-9155-e73f73b59817','Кукушка','1990-01-01',E'Песен еще ненаписанных, сколько?\nСкажи, кукушка, пропой.\nВ городе мне жить или на выселках,\nКамнем лежать или гореть звездой?\nЗвездой.\n\nСолнце мое – взгляни на меня,\nМоя ладонь превратилась в кулак,\nИ если есть порох – дай огня.\nВот так...\n\nКто пойдет по следу одинокому?\nСильные да смелые\nГоловы сложили в поле в бою.\nМало кто остался в светлой памяти,\nВ трезвом уме да с твердой рукой в строю,\nВ строю.\n\nСолнце мое – взгляни на меня,\nМоя ладонь превратилась в кулак,\nИ если есть порох – дай огня.\nВот так...\n\nГде же ты теперь, воля вольная?\nС кем же ты сейчас\nЛасковый рассвет встречаешь? Ответь.\nХорошо с тобой, да плохо без тебя,\nГолову да плечи терпеливые под плеть,\nПод плеть.\n\nСолнце мое – взгляни на меня,\nМоя ладонь превратилась в кулак,\nИ если есть порох – дай огня.\nВот так...','https://youtu.be/yNp9SBW4xTA'),
('3a7c2d18-29fb-4a7e-82d8-2800c510ec10','5176a448-10b9-465f-8e78-bb00502f9055','Прогулки по воде','1993-01-01',E'С пpичала pыбачил апостол Андpей\nА Спаситель ходил по воде\nИ Андpей доставал из воды пескаpей\nА Спаситель погибших людей\nИ Андpей закpичал: я покину пpичал\nЕсли ты мне откpоешь секpет\nИ Спаситель ответил: Спокойно Андpей\nНикакого секpета здесь нет\nВидишь там на гоpе возвышается кpест\nПод ним десяток солдат повиси-ка на нем\nА когда надоест возвpащайся назад\nГулять по воде\nГулять по воде\nГулять по воде со мной\n\n\nHо учитель на касках блистают pога\nЧеpный воpон кpужит над кpестом\nОбъясни мне сейчас пожалей дуpака\nА pаспятье оставь на потом\nОнемел Спаситель и топнул в сеpдцах\nПо водной глади ногой\nТы и веpно дуpак и Андpей в слезах\nПобpел с пескаpями домой\nВидишь там на гоpе возвышается кpест\nПод ним десяток солдат повиси-ка на нем\nА когда надоест возвpащайся назад\nГулять по воде\nГулять по воде\nГулять по воде со мной\nВидишь там на гоpе возвышается кpест\nПод ним десяток солдат повиси-ка на нем\nА когда надоест возвpащайся назад\nГулять по воде\nГулять по воде\nГулять по воде со мной','https://youtu.be/vQWUcWn-OJc'),
('96375e8f-c402-4548-84c8-a3c2e1991a16','cac0f25f-e305-43d6-9155-e73f73b59817','Группа крови','1988-01-01',E'Тёплое место\nНо улицы ждут oтпечатков наших ног\nЗвёздная пыль — на сапогах.\nМягкое кресло, клетчатый плед,\nНе нажатый вовремя курок.\nСолнечный день - в ослепительных снах.\n\nГруппа крови — на рукаве,\nМой порядковый номер — на рукаве,\nПожелай мне удачи в бою, пожелай мне:\nНе остаться в этой траве,\nНе остаться в этой траве.\nПожелай мне удачи, пожелай мне удачи!\n\nИ есть чем платить, но я не хочу\nПобеды любой ценой.\nЯ никому не хочу ставить ногу на грудь.\nЯ хотел бы остаться с тобой,\nПросто остаться с тобой,\nНо высокая в небе звезда зовёт меня в путь.\n\nГруппа крови — на рукаве,\nМой порядковый номер — на рукаве,\nПожелай мне удачи в бою, пожелай мне:\nНе остаться в этой траве,\nНе остаться в этой траве.\nПожелай мне удачи, пожелай мне удачи','https://youtu.be/k9p4B6s0j4U'),
('3459ebd5-936d-402c-9a78-19f706cbbb17','cac0f25f-e305-43d6-9155-e73f73b59817','Перемен','1989-01-01',E'Вместо тепла — зелень стекла,\nВместо огня — дым,\nИз сетки календаря выхвачен день.\nКрасное солнце сгорает дотла,\nДень догорает с ним,\nНа пылающий город падает тень.\n\nПеремен! — требуют наши сердца.\nПеремен! — требуют наши глаза.\nВ нашем смехе и в наших слезах,\nИ в пульсации вен —\nПеремен!\nМы ждём перемен!\n\nЭлектрический свет продолжает наш день,\nИ коробка от спичек пуста,\nНо на кухне синим цветком горит газ.\nСигареты в руках, чай на столе —\nЭта схема проста,\nИ больше нет ничего, всё находится в нас.\n\nПеремен! — требуют наши сердца.\nПеремен! — требуют наши глаза.\nВ нашем смехе и в наших слезах,\nИ в пульсации вен —\nПеремен!\nМы ждём перемен!\n\nМы не можем похвастаться мудростью глаз\nИ умелыми жестами рук,\nНам не нужно всё это, чтобы друг друга понять.\nСигареты в руках, чай на столе —\nТак замыкается круг,\nИ вдруг нам становится страшно что-то менять.\n\nПеремен! — требуют наши сердца.\nПеремен! — требуют наши глаза.\nВ нашем смехе и в наших слезах,\nИ в пульсации вен —\nПеремен!\nМы ждём перемен!','https://youtu.be/jyorQevSPI0'),
('c748cc8f-4843-4c09-aa19-dd652675043d','cac0f25f-e305-43d6-9155-e73f73b59817','В наших глазах','1988-01-01',E'Постой, не уходи!\nМы ждали лета - пришла зима.\nМы заходили в дома,\nНо в домах шел снег.\nМы ждали завтрашний день,\nКаждый день ждали завтрашний день.\nМы прячем глаза за шторами век.\n\nВ наших глазах крики Вперед!\nВ наших глазах окрики Стой!\nВ наших глазах рождение дня\nИ смерть огня.\nВ наших глазах звездная ночь,\nВ наших глазах потерянный рай,\nВ наших глазах закрытая дверь.\nЧто тебе нужно? Выбирай!\n\nМы хотели пить, не было воды.\nМы хотели света, не было звезды.\nМы выходили под дождь\nИ пили воду из луж.\nМы хотели песен, не было слов.\nМы хотели спать, не было снов.\nМы носили траур, оркестр играл туш...\n\nВ наших глазах крики Вперед!\nВ наших глазах окрики Стой!\nВ наших глазах рождение дня\nИ смерть огня.\nВ наших глазах звездная ночь,\nВ наших глазах потерянный рай,\nВ наших глазах закрытая дверь.\nЧто тебе нужно? Выбирай!','https://youtu.be/jyorQevSPI0'),
('63318518-828f-49d4-b93b-a7e9212222af','cac0f25f-e305-43d6-9155-e73f73b59817','Электричка','1982-01-01',E'Я вчера слишком поздно лёг, сегодня рано встал,\nЯ вчера слишком поздно лёг, я почти не спал.\nМне, наверно, с утра нужно было пойти к врачу,\nА теперь электричка везёт меня туда, куда я не хочу.\n\nЭлектричка везёт меня туда, куда я не хочу\n\nВ тамбуре холодно, и в то же время как-то тепло,\nВ тамбуре накурено, и в то же время как-то свежо.\nПочему я молчу, почему не кричу? Молчу.\n\nЭлектричка везёт меня туда, куда я не хочу','https://youtu.be/4lhxm6_aqAA'),
('a5bb1db4-bf4d-432b-9512-ea897c4c074a','cac0f25f-e305-43d6-9155-e73f73b59817','Мама, Мы Все Тяжело Больны','1988-01-01',E'Зерна упали в землю, зерна просят дождя.\nИм нужен дождь.\nРазрежь мою грудь, посмотри мне внутрь,\nТы увидишь, там все горит огнем.\nЧерез день будет поздно, через час будет поздно,\nЧерез миг будет уже не встать.\nЕсли к дверям не подходят ключи, вышиби двери плечом.\n\nМама, мы все тяжело больны...\nМама, я знаю, мы все сошли с ума...\n\nСталь между пальцев, сжатый кулак.\nУдар выше кисти, терзающий плоть,\nНо вместо крови в жилах застыл яд, медленный яд.\nРазрушенный мир, разбитые лбы, разломанный надвое хлеб.\nИ вот кто-то плачет, а кто-то молчит,\nА кто-то так рад, кто-то так рад..\n\nМама, мы все тяжело больны...\nМама, я знаю, мы все сошли с ума...\n\nТы должен быть сильным, ты должен уметь сказать:\nРуки прочь, прочь от меня!\nТы должен быть сильным, иначе зачем тебе быть.\nЧто будет стоить тысячи слов,\nКогда важна будет крепость руки?\nИ вот ты стоишь на берегу и думаешь: Плыть или не плыть?\n\nМама, мы все тяжело больны...\nМама, я знаю, мы все сошли с ума...','https://youtu.be/pLE2ngmVqZY'),
('6b150eac-3729-49ac-9aeb-48d9c9e983de','cac0f25f-e305-43d6-9155-e73f73b59817','Пачка сигарет','1989-01-01',E'Я сижу и смотрю в чужое небо из чужого окна\nИ не вижу ни одной знакомой звезды.\nЯ ходил по всем дорогам и туда, и сюда,\nОбернулся — и не смог разглядеть следы.\n\nНо если есть в кармане пачка сигарет,\nЗначит все не так уж плохо на сегодняшний день.\nИ билет на самолет с серебристым крылом,\nЧто, взлетая, оставляет земле лишь тень.\n\nИ никто не хотел быть виноватым без вина,\nИ никто не хотел руками жар загребать,\nА без музыки и на миру смерть не красна,\nА без музыки не хочется пропадать.\n\nНо если есть в кармане пачка сигарет,\nЗначит все не так уж плохо на сегодняшний день.\nИ билет на самолет с серебристым крылом,\nЧто, взлетая, оставляет земле лишь тень.','https://youtu.be/ciyB5a189_o'),
('e6670633-7396-424c-98db-fe2a0b63ac5a','33ce21ac-1875-4420-9e95-e422922f5e6d','У шамана три руки','2005-01-01',E'У шамана три руки\nИ крыло из-за плеча\nОт дыхания его\nРазгорается свеча\nИ порою сам себя\nСам себя не узнает\nА распахнута душа\nНадрывается, поет\n\nУ шамана три руки\nМир вокруг, как темный зал\nНа ладонях золотых\nНарисованы глаза\nВидит розовый рассвет\nПрежде солнца самого\nА казалось, будто спит\nИ не знает ничего\n\nУ шамана три руки.\nСад в рубиновых лучах.\nОт дыхания его\nРазгорается... Разгорается...\nРазгорается...\n\nРазгорается свеча...\nРазгорается свеча.\n\nУ шамана три...\nУ шамана три...\nУ шамана три...','https://youtu.be/Xhsl6j6xOEM'),
('1a8b1880-e3aa-4c60-823e-25e632642276','33ce21ac-1875-4420-9e95-e422922f5e6d','Там, на самом, на краю Земли','1991-01-01',E'Там, на самом на краю Земли\nВ небывалой голубой дали\nВнемля звукам небывалых слов,\nСладко-сладко замирает кровь.\n\nТам ветра летят, касаясь звeзд\nТам деревья не боятся гроз\nОкеаном бредят корабли\nТам, на самом, на краю Земли.\n\nЧто ж ты, сердце, рвeшься из груди?\nПогоди немного, погоди\nЧистый голос в небесах поeт\nСветлый полдень над Землeй встаeт.','https://youtu.be/YIdEbLa6L6s'),
('f8a11d33-cfc7-4dab-8e93-d3d35397b086','33ce21ac-1875-4420-9e95-e422922f5e6d','Из мышеловки','2007-01-01',E'Ещё грязь не смыта с кожи,\nТолько страха больше нет.\nПотому затих прохожий,\nУдивленно смотрит вслед.\n\nА движения неловки,\nБудто бы из мышеловки,\nБудто бы из мышеловки,\nТолько вырвалась она.\n\nА движения неловки,\nБудто бы из мышеловки,\nБудто бы из мышеловки,\nТолько вырвалась она.\n\nА в груди попеременно,\nБыл то пепел, то алмаз.\nТо лукаво то надменно,\nНочь прищуривала глаз.\n\nА движения неловки,\nБудто бы из мышеловки,\nБудто бы из мышеловки,\nТолько вырвалась она.\n\nЗа три слова до паденья,\nЗа второе за рожденье,\nЗа второе за рожденье,\nПей до полночи одна.\n\nИ уже не искалечит,\nСмех лощенных дураков.\nВедь запрыгнул ей на плечи,\nМокрый ангел с облаков.\n\nА движения...\n\nА движения неловки,\nБудто бы из мышеловки,\nБудто бы из мышеловки,\nТолько вырвалась она.','https://youtu.be/bS5ZXt-PVlY'),
('6b1b627f-1609-488c-9c38-c87a3b70d33d','5176a448-10b9-465f-8e78-bb00502f9055','Крылья','1995-01-01',E'Ты снимаешь вечернее платье,\nСтоя лицом к стене.\nИ я вижу свежие шрамы\nНа гладкой, как бархат, спине.\n\nМне хочется плакать от боли\nИли забыться во сне,\nГде твои крылья,\nКоторые так нравились мне?\n\nГде твои крылья,\nКоторые нравились мне?\nГде твои крылья,\nКоторые нравились мне?\n\nРаньше у нас было время,\nТеперь у нас есть дела —\nДоказывать, что сильный жрет слабых,\nДоказывать, что сажа бела.\n\nМы все потеряли что-то\nНа этой безумной войне.\nКстати, где твои крылья,\nКоторые нравились мне?\n\nГде твои крылья,\nКоторые нравились мне?\nГде твои крылья,\nКоторые нравились мне?\n\nЯ не спрашиваю, сколько у тебя денег,\nНе спрашиваю сколько мужей.\nЯ вижу — ты боишься открытых окон\nИ верхних этажей.\n\nИ если завтра начнется пожар\nИ все здание будет в огне,\nМы погибнем без этих крыльев,\nКоторые нравились мне.','https://youtu.be/ouKj-7bc-ZQ'),
('65347c0d-2a7f-4bf8-b892-ae7f012c417a','5176a448-10b9-465f-8e78-bb00502f9055','Музыка','1994-01-01',E'Звуки пойманные слухом\nВы сроднитесь с плотным духом\nВы вплетаете свободно\nВ пульс творенья что угодно\nСердца нашего биенья\nВы живые отраженья\n\nЗвук обступает и слева и справа\nСмутным таинственным войском вдали\nПеной ложится у кромки земли\nИ до утра до скончания ночи\nЗвук обступает и слева и справа\nБесится бегает дико хохочетс\nСмутным таинственным войском вдали\n\nСтрелами света кружит над тобой\nСхватка оркестра с волшебной трубой\nчтобы и завтра безвестная сила\nВсе повторяла и все возносила Все повторяла и все возносила\nСловно и радость твоя и тоска\nСтрелами света кружит над тобой\n\nМоре бурное оркестра\nВот тогда душа моя\nСнова музыка поманит\nЗвуки пойманные слухом\nВы сроднитесь с плотным духом\nВы летаете свободно\n\nЗа крылатым вашим дымом\nЯ сыщу владенья ваши','https://youtu.be/0aTzMbrAW20');
--add_group
CREATE OR REPLACE FUNCTION public.add_group(
	request jsonb DEFAULT '{}'::jsonb)
    RETURNS json
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
AS $BODY$
DECLARE response json;
BEGIN
	WITH t AS(
		INSERT INTO public.groups
		VALUES (
			gen_random_uuid(),
			(add_group.request ->> 'name')::character varying
		)
		RETURNING group_id,"name"
	)
	SELECT row_to_json(t) INTO response FROM t;
	RETURN response;
END;
$BODY$;
ALTER FUNCTION public.add_group(jsonb)
    OWNER TO postgres;
--add_song
CREATE OR REPLACE FUNCTION public.add_song(
	request jsonb DEFAULT '{}'::jsonb)
    RETURNS json
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
AS $BODY$
DECLARE response json;
BEGIN
	WITH t AS (
		SELECT group_id, "name"
		FROM public.groups
		WHERE LOWER(groups.name)=LOWER((add_song.request ->>'group')::character varying)
	),
	t2 AS (
		INSERT INTO public.songs
		VALUES (
			gen_random_uuid(),
			(SELECT group_id FROM t),
			(add_song.request ->> 'name')::character varying,
			(add_song.request ->> 'release_date')::date,
			(add_song.request ->> 'text')::character varying,
			(add_song.request ->> 'link')::character varying
		)
		RETURNING song_id,"name",release_date,"text","link"
	)
	SELECT row_to_json(t3) INTO response FROM (SELECT t2.*, row_to_json(t) AS group  FROM t,t2) t3;
	RETURN response;
END;
$BODY$;
ALTER FUNCTION public.add_song(jsonb)
    OWNER TO postgres;
--get_group
CREATE OR REPLACE FUNCTION public.get_group(
	request jsonb DEFAULT '{}'::jsonb)
    RETURNS SETOF json 
    LANGUAGE 'plpgsql'
    COST 100
    STABLE PARALLEL UNSAFE
    ROWS 1000
AS $BODY$
BEGIN
	RETURN QUERY
	WITH t AS (
		SELECT group_id,"name" 
		FROM public.groups
		WHERE groups.group_id=(get_group.request ->> 'group_id')::uuid
	)
	SELECT row_to_json(t) 
	FROM t;
END;
$BODY$;
ALTER FUNCTION public.get_group(jsonb)
    OWNER TO postgres;
--get_groups
CREATE OR REPLACE FUNCTION public.get_groups(
	request jsonb DEFAULT '{}'::jsonb)
    RETURNS json
    LANGUAGE 'plpgsql'
    COST 100
    STABLE PARALLEL UNSAFE
AS $BODY$
DECLARE response json;
BEGIN
	WITH t AS (
		SELECT group_id, "name"
		FROM public.groups
		WHERE 
			(get_groups.request ->>'name') IS null OR 
			LOWER(groups.name) LIKE '%' || LOWER((get_groups.request ->>'name')::character varying) || '%'
	)
	SELECT json_build_object( 
		'data',
		COALESCE(
			(SELECT json_agg(row_to_json(t2)) FROM (SELECT * FROM t OFFSET (get_groups.request ->>'offset')::integer LIMIT (get_groups.request ->>'limit')::integer) t2), 
			'[]'::json),
		'all_record_count', 
		(SELECT count(*) FROM t)
	) INTO response;
RETURN response;
END;
$BODY$;
ALTER FUNCTION public.get_groups(jsonb)
    OWNER TO postgres;
--get_song
CREATE OR REPLACE FUNCTION public.get_song(
	request jsonb DEFAULT '{}'::jsonb)
    RETURNS SETOF json 
    LANGUAGE 'plpgsql'
    COST 100
    STABLE PARALLEL UNSAFE
    ROWS 1000
AS $BODY$
BEGIN
	RETURN QUERY
	WITH t AS (
		SELECT 
			_s.song_id,
			_s.name,
			_s.release_date,
			_s.text,
			_s.link,
			row_to_json(_g) AS "group"
		FROM public.songs _s
		JOIN public.groups _g
		ON _g.group_id=_s.group_id
		WHERE _s.song_id=(get_song.request ->>'song_id')::uuid
	)
	SELECT row_to_json(t) 
	FROM t;
END;
$BODY$;
ALTER FUNCTION public.get_song(jsonb)
    OWNER TO postgres;
--get_song_text
CREATE OR REPLACE FUNCTION public.get_song_text(
	request jsonb DEFAULT '{}'::json)
    RETURNS SETOF json 
    LANGUAGE 'plpgsql'
    COST 100
    STABLE PARALLEL UNSAFE
    ROWS 1000
AS $BODY$
BEGIN
	RETURN QUERY
	WITH t AS (
		SELECT *
		FROM public.songs
		WHERE songs.song_id=(get_song_text.request ->> 'song_id')::uuid
	),
	t2 AS (
		SELECT ROW_NUMBER() OVER() AS "index",	"text"
		FROM (SELECT string_to_table(t.text,E'\n\n') AS "text" FROM t)
	)
	SELECT json_build_object(
		'data',
		COALESCE(
			(SELECT json_agg(row_to_json(t3)) FROM (SELECT * FROM t2 OFFSET (get_song_text.request ->>'offset')::integer LIMIT (get_song_text.request ->>'limit')::integer) t3),
			'[]'::json),
		'all_record_count',
		(SELECT COUNT(*) FROM t2)) 
	FROM t;
END;
$BODY$;
ALTER FUNCTION public.get_song_text(jsonb)
    OWNER TO postgres;
--get_songs
CREATE OR REPLACE FUNCTION public.get_songs(
	request jsonb DEFAULT '{}'::json)
    RETURNS json
    LANGUAGE 'plpgsql'
    COST 100
    STABLE PARALLEL UNSAFE
AS $BODY$
DECLARE response json;
BEGIN
	WITH t AS (
		SELECT 
			_s.song_id,
			_s.name,
			_s.release_date,
			_s.text,
			_s.link,
			row_to_json(_g) AS "group"
		FROM public.songs _s
		JOIN public.groups _g
		ON _g.group_id=_s.group_id
		WHERE 
			((get_songs.request ->>'group') IS null OR LOWER(_g.name) LIKE '%' || LOWER((get_songs.request ->>'group')::character varying) || '%') AND
			((get_songs.request ->>'name') IS null OR LOWER(_s.name) LIKE '%' || LOWER((get_songs.request ->>'name')::character varying) || '%') AND
			((get_songs.request ->>'release_date') IS null OR _s.release_date = (get_songs.request ->>'release_date')::date) AND
			((get_songs.request ->>'text') IS null OR LOWER(_s.text) LIKE '%' || LOWER((get_songs.request ->>'text')::character varying) || '%') AND
			((get_songs.request ->>'link') IS null OR LOWER(_s.link) LIKE '%' || LOWER((get_songs.request ->>'link')::character varying) || '%') 
	)
	SELECT json_build_object( 
		'data',
		COALESCE(
			(SELECT json_agg(row_to_json(t2)) FROM (SELECT * FROM t OFFSET (get_songs.request ->>'offset')::integer LIMIT (get_songs.request ->>'limit')::integer) t2), 
			'[]'::json),
		'all_record_count', 
		(SELECT count(*) from t)
	) INTO response;
RETURN response;
END;
$BODY$;
ALTER FUNCTION public.get_songs(jsonb)
    OWNER TO postgres;
--remove_group
CREATE OR REPLACE FUNCTION public.remove_group(
	request jsonb DEFAULT '{}'::jsonb)
    RETURNS SETOF json 
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
    ROWS 1000
AS $BODY$
BEGIN
	RETURN QUERY
	WITH t AS (
		DELETE FROM public.groups
		WHERE groups.group_id=(remove_group.request ->>'group_id')::uuid
		RETURNING group_id
	)
	SELECT row_to_json(t) 
	FROM t;
END;
$BODY$;
ALTER FUNCTION public.remove_group(jsonb)
    OWNER TO postgres;
--remove_song
CREATE OR REPLACE FUNCTION public.remove_song(
	request jsonb DEFAULT '{}'::jsonb)
    RETURNS SETOF json 
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
    ROWS 1000
AS $BODY$
BEGIN
	RETURN QUERY
	WITH t AS (
		DELETE FROM public.songs
		WHERE songs.song_id=(remove_song.request ->>'song_id')::uuid
		RETURNING song_id
	)
	SELECT row_to_json(t)
	FROM t;
END;
$BODY$;
ALTER FUNCTION public.remove_song(jsonb)
    OWNER TO postgres;
--update_group
CREATE OR REPLACE FUNCTION public.update_group(
	request jsonb DEFAULT '{}'::jsonb)
    RETURNS SETOF json 
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
    ROWS 1000
AS $BODY$
BEGIN
	RETURN QUERY
	WITH t AS (
		UPDATE public.groups
		SET	"name"=COALESCE((update_group.request ->>'name')::character varying,"name")
		WHERE groups.group_id=(update_group.request ->>'group_id')::uuid
		RETURNING group_id
	)
	SELECT row_to_json(t) 
	FROM t;
END;
$BODY$;
ALTER FUNCTION public.update_group(jsonb)
    OWNER TO postgres;
--update_song
CREATE OR REPLACE FUNCTION public.update_song(
	request jsonb DEFAULT '{}'::jsonb)
    RETURNS SETOF json 
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
    ROWS 1000
AS $BODY$
BEGIN
	WITH t AS (
		UPDATE public.songs
		SET	
			"group_id"=COALESCE((update_song.request ->>'group_id')::uuid,group_id),
			"name"=COALESCE((update_song.request ->>'name')::character varying,"name"),
			release_date= CASE WHEN (update_song.request ->>'set_release_date')::boolean=true OR  (update_song.request ->>'release_date') IS NOT null THEN (update_song.request ->>'release_date')::date ELSE release_date END,
			"text"= CASE WHEN (update_song.request ->>'set_text')::boolean=true OR (update_song.request ->>'text') IS NOT null THEN (update_song.request ->>'text')::character varying ELSE songs.text END,
			"link"= CASE WHEN (update_song.request ->>'set_link')::boolean=true OR (update_song.request ->>'link') IS NOT null THEN (update_song.request ->>'link')::character varying ELSE songs.link END
		WHERE songs.song_id=(update_song.request ->>'song_id')::uuid
		RETURNING song_id
	)
	SELECT row_to_json(t)
	FROM t;
END;
$BODY$;
ALTER FUNCTION public.update_song(jsonb)
    OWNER TO postgres;
-- +goose StatementEnd
-- +goose Down
DROP DATABASE song_library;