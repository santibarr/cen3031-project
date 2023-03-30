package scrape

import (
	"testing"
)

func TestInWebsiteMiamiDade(t *testing.T) {
	catsItem := "{TARZAN (A2246622) Male Domestic Shorthair 1 year, 10 months old /image/538520640} {GUMDROP (A2332730) Female Domestic Shorthair 3 years old /image/537097077} {LITTLE SUSIE (A2334975) Female Domestic Shorthair 1 year, 1 month old /image/534375228} {GWEN (A2336604) Female Domestic Shorthair 1 year, 1 month old /image/537097078} {STEVIE (A2336609) Female Domestic Shorthair 1 year, 1 month old /image/537097079} {MJ (A2336611) Female Domestic Shorthair 1 year, 1 month old /image/537097080} {MILES (A2336613) Male Domestic Shorthair 1 year, 1 month old /image/534378996} {BOBBY (A2358734) Male Domestic Shorthair 10 months old /image/537097082} {LEILA (A2364680) Female Domestic Shorthair 1 year, 9 months old /image/526531546} {DASH (A2367295) Male Domestic Shorthair 5 years old /image/522591569} {MATTHEW (A2392264) Male Domestic Shorthair mix 7 months old /image/537097086} {TUXY (A2405402) Male Domestic Shorthair 4 years old /image/531774571} {ABBY (A2405406) Female Domestic Shorthair 4 years old /image/531356236} {ALEX (A2405409) Male Domestic Shorthair 4 years old /image/537097092} {AVA (A2405410) Female Domestic Shorthair 4 years old /image/531774572} {TODO (A2409809) Male Domestic Shorthair 10 months old /image/538878725} {BARBIE (A2424474) Female Domestic Shorthair 2 years, 2 months old /image/538013467} {TRIFECTA (A2426544) Male Domestic Shorthair 6 years old /image/535572014} {BERNIE (A2426910) Male Domestic Shorthair 1 year, 7 months old /image/539163166} {LEAH (A2429498) Female Domestic Shorthair 8 years old /image/536131543} {ZORRO (A2429753) Male Domestic Shorthair 4 years old /image/537120718} {HEINZ (A2430828) Male Domestic Shorthair 8 years old /image/536047759} {TIGER (A2431195) Male Domestic Shorthair 5 months old /image/539217215} {KIMBA (A2431199) Male Domestic Shorthair 1 year, 2 months old /image/539000487} {TIGRE (A2431200) Male Domestic Shorthair 1 year, 2 months old /image/538993760} {LUNA (A2432446) Female Domestic Shorthair 2 years, 2 months old /image/539243988} {LILA (A2433349) Female Domestic Shorthair 3 months old /image/538533620} {LAYLA (A2433353) Female Domestic Shorthair 3 months old /image/538533622} {MARY (A2434373) Female Domestic Shorthair 1 year, 1 month old /image/537097099} {NIKO (A2436572) Male Domestic Shorthair 4 months old /image/537971652} "
	expected := catsItem
	actual := miamiDadeString()
	if expected != actual {
		t.Errorf("Expected %s do not match actual %s", expected, actual)
	}
}
func TestInWebsiteLakeCounty(t *testing.T) {
	catsItem := "{Leo Neutered Male On Website Shelter staff think I am about 3 years old. /image/538940149} {Sage Neutered Male On Website Shelter staff think I am about 4 years old. /image/538233986} {Monkey Soda Neutered Male On Website Shelter staff think I am about 7 years old. /image/539143365} {Jimin Neutered Male On Website Shelter staff think I am about 1 year and 6 months old. /image/538877647} {Hope Neutered Male On Website Shelter staff think I am about 1 year and 6 months old. /image/538877648} {Jk Neutered Male On Website Shelter staff think I am about 1 year and 6 months old. /image/538877650} {Ivory Spayed Female On Website Shelter staff think I am about 8 years old. /image/539164084} "
	expected := catsItem
	actual := lakeCountyString()
	if expected != actual {
		t.Errorf("Expected %s do not match actual %s", expected, actual)
	}
}
func TestInWebsitePeggy(t *testing.T) {
	catsItem := "{Babe Girl On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0433576&LOCATION=PBHS} {Bear Bear On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0432381&LOCATION=PBHS} {Beau On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0266224&LOCATION=PBHS} {Bella On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0204866&LOCATION=PBHS} {Benji On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0373433&LOCATION=PBHS} {Bigs On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0433579&LOCATION=PBHS} {Boo Boo On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0426410&LOCATION=PBHS} {Boots On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0434186&LOCATION=PBHS} {Briggs On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0433484&LOCATION=PBHS} {Bucky On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0433582&LOCATION=PBHS} {Cabana On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0414568&LOCATION=PBHS} {Carl On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0428889&LOCATION=PBHS} {Carrie On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0432426&LOCATION=PBHS} {Cookie On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0431219&LOCATION=PBHS} {Crystal Blue On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0330344&LOCATION=PBHS} {Daisy On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0433672&LOCATION=PBHS} {Devi On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0431800&LOCATION=PBHS} {Dexter On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0431603&LOCATION=PBHS} {Emma On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0347064&LOCATION=PBHS} {Ettore On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0262356&LOCATION=PBHS} {Ezra On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0431014&LOCATION=PBHS} {Fruina On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0433627&LOCATION=PBHS} {Gunther On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0388231&LOCATION=PBHS} {Jellybean On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0433626&LOCATION=PBHS} {Jimmy On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0434184&LOCATION=PBHS} {Khaleesi On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0216863&LOCATION=PBHS} {Kitie On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0364748&LOCATION=PBHS} {Kitty On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0319815&LOCATION=PBHS} {Kobi On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0216865&LOCATION=PBHS} {Koko On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0431360&LOCATION=PBHS} {Kyrie On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0432978&LOCATION=PBHS} {Maple On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0348776&LOCATION=PBHS} {Maynard On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0320103&LOCATION=PBHS} {Mazie On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0433552&LOCATION=PBHS} {Mika On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0433314&LOCATION=PBHS} {Minnie On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0433343&LOCATION=PBHS} {Moon On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0433683&LOCATION=PBHS} {Naomi On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0298210&LOCATION=PBHS} {Nike On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0431836&LOCATION=PBHS} {Pikachu On Website On Website On Website https://petharbor.com/get_image.asp?RES=thumb&ID=A0433874&LOCATION=PBHS} "
	expected := catsItem
	actual := peggyString()
	if expected != actual {
		t.Errorf("Expected %s do not match actual %s", expected, actual)
	}
}

func TestInWebsiteKeyWest(t *testing.T) {
	catsItem := "{Magoo Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/0c901efd-827d-4423-ae9e-274dbc2a4d3c.jpg} {Maisey Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/9b0e4c40-e238-4917-b93c-7dd2b6e7e607.jpg} {Willie Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/9e3df4be-1605-41d4-9fb4-a0cfe8669cbb.jpg} {Shadow Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/afa1bd2e-eae8-473d-aa9d-b4950b7b084d.jpg} {Rosie Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/390595d7-a397-4006-9450-b8c6d42f07a0.jpg} {Ally Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/9a1c3456-6048-4642-9681-aa4d8b7bcc1b.jpg} {Zuna Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/c72897ca-fb6c-43e0-82e9-639db950bf24.jpg} {Dreamsicle Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/4756eda7-7dbd-412b-9c9e-fa3d964321b7.jpg} {Dimple Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/e86c00ce-c6f2-46fd-bb93-d91fb195f8bd.jpg} {Ninja Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/36ddf3b6-fcec-4079-987f-23c180b0e7ca.jpg} {Kinsley Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/fc40398a-6390-4863-86ae-d47772a0b73b.jpg} {Tiny Tina Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/152005a3-46bd-49e3-9395-9e9a441c0156.jpg} {Espresso Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/f6697b78-3c8d-4c31-9845-8642cfd98c2e.jpg} {Meech Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/64522aa3-8bd4-49ad-b5c2-beeb3f0f3185.jpg} {Louise Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/0d2fc55c-f849-47b9-9e7f-559f78823f78.jpg} {Boone Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/af53cad1-491e-45bf-8538-a52c2ad94cbe.jpg} {Brittney Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/d9a25ba3-68c3-49e5-9d54-d7d0906f39d9.jpg} {Billy Joe Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/ad4d086f-1276-481e-bae4-6e127b0467f0.jpg} {Penny Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/1a850656-7b91-46b6-9061-5ee942a13a08.jpg} {Shandor Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/b4e3bdc4-fae7-4f01-b541-52845c944eed.jpg} {Sophia Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/d5502461-9780-42a1-8cc4-f7d5c265bcc4.jpg} {Kolie Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/69bccc10-d941-4e96-9edc-38bd16124067.jpg} {Spritz Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/b875aeea-8638-4d6c-b18d-9b125576a9b8.jpg} {Crackle Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/349a42ab-d1f6-4ba4-9c4c-097f06960545.jpg} {Double Bubble Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/0d4e3ebe-53fd-450e-8b30-f92a75218b07.jpg} {Maya Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/20242e41-a2b1-4a66-b5a1-b066fa23358b.jpg} {Herman Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/282eb3a6-7077-4a15-ae75-9ec128245252.jpg} {Ron Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/8f005689-f582-4fcf-a755-936739995658.jpg} {Dobby Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/d49170bd-9f82-47d9-b139-76b41127b90b.jpg} {Gobi Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/8490365f-9209-4755-b2a6-6d73e9c3c907.jpg} {Emmanuel Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/85596ac7-c26b-4748-afa6-12cd42d955c3.jpg} {Sonny Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/bed62086-aede-47a6-82b8-59f9e95b5d03.jpg} {Jason Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/dd2239df-420a-4a0d-8100-7226c8b9a7a8.jpg} {Klaus Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/05d835cf-c06a-4466-b363-a73d99379df0.jpg} {Buddha Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/b074bead-119d-418f-bfed-8f69e1cfcaa9.jpg} {Christmas Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/45c72ace-7a80-4b27-a641-d5841ef9a620.jpg} {Wheaties Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/3b9b92ea-01b6-4bf4-b8be-6b9b63e6c378.jpg} {Kimchi Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/4a0cd2af-4af9-4dca-bf5c-4f82e951c7d2.jpg} {Dashi Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/35732d55-54de-4116-a936-8b41fd9e7849.jpg} {Udon Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/d7f792f4-aebc-4846-896c-5d743527293b.jpg} {Falafel Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/905965dc-ae20-4a65-a53c-458673e17888.jpg} {Gelato Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/3e712e32-c2bd-4a8c-8d92-b9e22537a62c.jpg} {Tahini Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/f089a915-8d51-43e6-a62c-365b02d623f6.jpg} {Tortellini Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/0be71903-22b0-42d6-95c1-b6ec72331611.jpg} {Lemon Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/568e854d-9bf1-4ff5-a732-5768b1a8c60c.jpg} {Poppy Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/bb69fec6-1b2f-419f-8b91-56a6a78ae247.jpg} {Kerrigan Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/90bcb7e5-ed8e-4218-a9fc-3a6610d8542f.jpg} {Ed Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/a31e2dcb-ea27-4372-809b-cd39d603fbd8.jpg} {Crepe Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/720736da-d23c-4238-99d9-a70ec664f4be.jpg} {Pita Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/37c16c69-c785-49d4-9759-a739b80c4220.jpg} {Nox Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/7c2a89e8-9f3d-42cf-941c-6bc1ee69bddf.jpg} {Nash Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/6292f432-f221-45fb-91c3-af46a7d5e7c8.jpg} {Taos Male Turkish Angora/Mix On Website https://g.petango.com/photos/370/d3a02238-aaf3-48db-89d2-baaef96748a2.jpg} {Mango Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/12ad9a8d-8ac7-4aaa-93d3-516af04ac5e6.jpg} {Tyler Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/f35c3d8d-6dca-400d-8bfe-a33b79e89451.jpg} {Snap Female Domestic Medium Hair/Mix On Website https://g.petango.com/photos/370/877e2b6f-d385-48c4-acab-23339c1f7a9d.jpg} {Pop Female Domestic Medium Hair/Mix On Website https://g.petango.com/photos/370/e1648f65-0e98-4e21-a645-5d992320b689.jpg} {Rabbit Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/aac355d7-674e-4301-aa46-5f5d8aca8398.jpg} {Velvet Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/832cd522-45fb-4dda-8bf6-f2f852b07f7c.jpg} {Ornament Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/2b581d6c-97f9-4649-8ecb-a7ce26634972.jpg} {Garland Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/abf5083c-e797-4eb4-b68f-87eb628318d5.jpg} {Candy Star Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/937217b9-bb5d-48b5-9c4d-62f09e23f30d.jpg} {Mello Yello Male Turkish Angora/Mix On Website https://g.petango.com/photos/370/b0cd72ac-7830-4cb6-a814-d6212277655b.jpg} {Lorelai Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/8bb4c56f-d226-4a3a-a57f-fd3a0c8c230b.jpg} {Sebastian Male Domestic Medium Hair/Mix On Website https://g.petango.com/photos/370/cf31901f-fd39-4132-9671-7cfff57cbbbb.jpg} {Zeke Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/ab2cef1d-fc50-4cbc-9237-64151212f268.jpg} {Stetson Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/488cf7a4-0986-426d-a034-53a7b4c8450a.jpg} {Charrito Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/0e9639d9-6ade-40cd-8a49-22976631bb9a.jpg} {Jaguar Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/3423b3e2-47bd-45c3-819a-5d042e393748.jpg} {Star Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/eab8d8c7-e23b-4233-ac7a-59b0ac342744.jpg} {Sailboat Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/43b59a54-edc0-45af-b4a9-70117fac50c3.jpg} {Tiger Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/a7ec3247-83c3-493a-a7b4-888b089af1ff.jpg} {Twinklenose Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/60119ce4-6798-4919-889d-b152dc7481fd.jpg} {Macie Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/6165cdbd-71b9-43a4-a60d-7a1f05e67139.jpg} {Gremlin Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/7ae20838-cfda-4ff2-8a70-a10fc420638d.jpg} {Davidson Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/fb39c2e0-e5ab-4a77-8611-c5647a331eaa.jpg} {Fruit Stripe Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/44b83c2c-3266-4866-824d-e8f98e58b3f5.jpg} {Mr. Chunks Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/60e3467d-3918-4647-aa67-5e240fcad8ec.jpg} {Binx Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/e0a114d1-6e61-4fb0-828f-76e8018279ef.jpg} {PeePaw Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/485cad5d-0676-4ea3-b66f-a6509901a80f.jpg} {Aries Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/008a5538-07a3-43f7-b1fa-041a683bd6c6.jpg} "
	expected := catsItem
	actual := keyWestString()
	if expected != actual {
		t.Errorf("Expected %s do not match actual %s", expected, actual)
	}
}
func TestInWebsiteMarathon(t *testing.T) {
	catsItem := "{Candi Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/94aae0fb-60b6-40d4-af5e-b1f93049f1b5.jpg} {Emily Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/763df0c2-557c-46a3-a21e-f8b99e29e80b.jpg} {Skyy Female American Shorthair/Mix On Website https://g.petango.com/photos/370/418e2721-4c35-45ef-a61b-c38ca0bb955f.jpg} {Whopper Female Domestic Medium Hair/Mix On Website https://g.petango.com/photos/370/d4d005d1-748e-4fca-9145-d22d3a5dfd69.jpg} {Cornbread Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/5c1e5aca-6a96-4a94-bf69-03db834dedce.jpg} {Sunshine Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/1b9f908e-e095-4e41-9a3a-a38168355058.jpg} {Jethro Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/11329bef-30f3-42db-bdd7-23a3a0a73273.jpg} {Blondi Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/14d2e966-98a0-4771-a0ae-357ee23f25a1.jpg} {Baby Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/25eedb98-0ca6-44d4-8389-d30d8e966f76.jpg} {Cowboy (Golden Paw) Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/5da00e92-1f34-42a1-9ccc-a4997dbc8193.jpg} {Pugsley Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/5ff90e1d-140a-48f2-ae9d-16e39d50fb17.jpg} {Wednesday Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/e694a6f9-ca76-4896-85d2-f1bbf0867999.jpg} {Aria Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/c405ddf6-8297-4904-8be3-cc99e22cca19.jpg} {Doodle Unknown Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/a273d761-b579-4ffe-b7d0-828627b9aeab.jpg} {Cinnamon Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/8b71564e-9082-4dcd-8f98-e26990d9b375.jpg} {Tux Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/5e98e59a-282e-400b-84e8-5d97cc3a279d.jpg} {Cali Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/46a54877-87fb-4535-84a4-4087f286de0d.jpg} {On Website On Website On Website On Website On Website} "
	expected := catsItem
	actual := marathonString()
	if expected != actual {
		t.Errorf("Expected %s do not match actual %s", expected, actual)
	}
}

func TestInWebMarathonAge(t *testing.T) {
	var catList []Cat
	catList = Marathon(catList)
	expected := "On Website"
	for i := 0; i < len(catList); i++ {
		actual := catList[i].Age
		if expected != actual {
			t.Errorf("Expected %s do not match actual %s", expected, actual)
		}
	}
	//iterate through the catItem list and check if the Age says On Website since the HTML extraction did not display the age

}
func TestInWebLakeCountyBreed(t *testing.T) {
	var catList []Cat
	catList = LakeCounty(catList)
	expected := "On Website"
	for i := 0; i < len(catList); i++ {
		actual := catList[i].Breed
		if expected != actual {
			t.Errorf("Expected %s do not match actual %s", expected, actual)
		}
	}
}
//--------------------------------------------------------Sprint 3 Tests----------------------------------------------------------------
func TestFeaturesMiamiDade(t *testing.T){
	var catList []Cat

	m := map[string]int{
		"friendly": 1,
		"playful": 2,
		"adorable": 3,
		"loving": 4,
		"affectionate": 5,
		"curious": 6,
		"energetic": 7,
		"loves attention": 8,
		"loves to cuddle": 9,
		"loves to play": 10,
		"loves to be pet": 11,
	}

	catList = MiamiDade(catList)
	for i := 0; i < len(catList); i++ {
		actual := catList[i].Feature
		if _, ok := m[actual]; !ok {
			t.Errorf("Expected features in given features do not match actual %s", actual)
		}

	}
}
func TestFeaturesLakeCounty(t *testing.T){
	var catList []Cat

	m := map[string]int{
		"friendly": 1,
		"playful": 2,
		"adorable": 3,
		"loving": 4,
		"affectionate": 5,
		"curious": 6,
		"energetic": 7,
		"loves attention": 8,
		"loves to cuddle": 9,
		"loves to play": 10,
		"loves to be pet": 11,
	}

	catList = LakeCounty(catList)
	for i := 0; i < len(catList); i++ {
		actual := catList[i].Feature
		if _, ok := m[actual]; !ok {
			t.Errorf("Expected features in given features do not match actual %s", actual)
		}

	}
}
func TestFeaturesPeggy(t *testing.T){
	var catList []Cat

	m := map[string]int{
		"friendly": 1,
		"playful": 2,
		"adorable": 3,
		"loving": 4,
		"affectionate": 5,
		"curious": 6,
		"energetic": 7,
		"loves attention": 8,
		"loves to cuddle": 9,
		"loves to play": 10,
		"loves to be pet": 11,
	}

	catList = Peggy(catList)
	for i := 0; i < len(catList); i++ {
		actual := catList[i].Feature
		if _, ok := m[actual]; !ok {
			t.Errorf("Expected features in given features do not match actual %s", actual)
		}

	}
}
