package scrape

import (
	"testing"
)

func TestInWebsiteMiamiDade(t *testing.T) {
	catsItem := "{MANDY (A1791429) Female Domestic Shorthair 7 years old /image/540531310} {JIMMY LEE (A1796744) Male Domestic Shorthair 7 years old /image/540531312} {CATHY (A2144328) Female Domestic Shorthair 3 years old /image/540396515} {TARZAN (A2246622) Male Domestic Shorthair 1 year, 11 months old /image/538520640} {TEDO (A2301229) Male Domestic Shorthair mix 1 year, 5 months old /image/540707615} {GUMDROP (A2332730) Female Domestic Shorthair 3 years old /image/539594833} {LITTLE SUSIE (A2334975) Female Domestic Shorthair 1 year, 2 months old /image/539594837} {GWEN (A2336604) Female Domestic Shorthair 1 year, 2 months old /image/539594838} {STEVIE (A2336609) Female Domestic Shorthair 1 year, 2 months old /image/539594839} {MJ (A2336611) Female Domestic Shorthair 1 year, 2 months old /image/537097080} {MILES (A2336613) Male Domestic Shorthair 1 year, 2 months old /image/534378996} {BOBBY (A2358734) Male Domestic Shorthair 11 months old /image/539594842} {DASH (A2367295) Male Domestic Shorthair 5 years old /image/522591569} {TUXY (A2405402) Male Domestic Shorthair 4 years old /image/539594850} {ABBY (A2405406) Female Domestic Shorthair 4 years old /image/531356236} {ALEX (A2405409) Male Domestic Shorthair 4 years old /image/539594851} {AVA (A2405410) Female Domestic Shorthair 4 years old /image/531774572} {NALA (A2406262) Female Domestic Shorthair 1 year, 2 months old /image/540385628} {TODO (A2409809) Male Domestic Shorthair 11 months old /image/539594853} {TRIFECTA (A2426544) Male Domestic Shorthair 6 years old /image/535572014} {BERNIE (A2426910) Male Domestic Shorthair 1 year, 8 months old /image/539283144} {HEINZ (A2430828) Male Domestic Shorthair 8 years old /image/536047759} {KIMBA (A2431199) Male Domestic Shorthair 1 year, 3 months old /image/539594583} {DAISY (A2437938) Female Domestic Shorthair 5 months old /image/539594862} {KEYLA (A2438036) Female Domestic Shorthair 2 years, 2 months old /image/537097100} {KASEY (A2438057) Male Domestic Shorthair 12 weeks old /Content/Images/No_pic_t.jpg} {KYLE (A2438060) Male Domestic Shorthair 12 weeks old /Content/Images/No_pic_t.jpg} {KANO (A2438062) Male Domestic Shorthair 12 weeks old /Content/Images/No_pic_t.jpg} {SPOT (A2442084) Male Domestic Shorthair 2 years, 2 months old /image/540351866} {LUCKY (A2442538) Female Domestic Shorthair 6 years old /image/540255402} "
	expected := catsItem
	actual := miamiDadeString()
	if expected != actual {
		t.Errorf("Expected %s do not match actual %s", expected, actual)
	}
}
func TestInWebsiteLakeCounty(t *testing.T) {
	catsItem := "{Smudge Neutered Male On Website Shelter staff think I am about 2 years and 2 months old. /image/540511618} {Shadow Neutered Male On Website Shelter staff think I am about 3 years old. /image/540681733} {Chowder Neutered Male On Website Shelter staff think I am about 2 years old. /image/540503455} {Jonah Neutered Male On Website Shelter staff think I am about 8 years old. /image/539931968} {Jubee Spayed Female On Website Shelter staff think I am about 3 years old. /image/539771451} {Evie Spayed Female On Website Shelter staff think I am about 3 years old. /image/540684569} {Rudy Neutered Male On Website Shelter staff think I am about 5 years old. /image/540683250} {Pixie Unaltered Female On Website Shelter staff think I am about 1 year old. /image/540503456} {A056602 Spayed Female On Website Shelter staff think I am about 2 years old. /image/540688298} {Snowflake Spayed Female On Website Shelter staff think I am about 5 years old. /image/540488875} {Angel Spayed Female On Website Shelter staff think I am about 3 years old. /image/540465778} {Atica Spayed Female On Website Shelter staff think I am about 3 years old. /image/540465779} {Boo Bear Neutered Male On Website Shelter staff think I am about 5 years old. /image/540465780} {Charlotte Spayed Female On Website Shelter staff think I am about 3 years old. /image/540465781} {Gayle Spayed Female On Website Shelter staff think I am about 3 years old. /image/540465782} "
	expected := catsItem
	actual := lakeCountyString()
	if expected != actual {
		t.Errorf("Expected %s do not match actual %s", expected, actual)
	}
}
func TestInWebsitePeggy(t *testing.T) {
	catsItem := " "
	expected := catsItem
	actual := peggyString()
	if expected != actual {
		t.Errorf("Expected %s do not match actual %s", expected, actual)
	}
}

func TestInWebsiteKeyWest(t *testing.T) {
	catsItem := "{Magoo Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/0c901efd-827d-4423-ae9e-274dbc2a4d3c.jpg} {Maisey Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/9b0e4c40-e238-4917-b93c-7dd2b6e7e607.jpg} {Willie Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/9e3df4be-1605-41d4-9fb4-a0cfe8669cbb.jpg} {Shadow Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/afa1bd2e-eae8-473d-aa9d-b4950b7b084d.jpg} {Rosie Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/390595d7-a397-4006-9450-b8c6d42f07a0.jpg} {Brad Pitt Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/f20bf4a6-e554-4270-b73f-51bc9a05efc6.jpg} {Ally Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/9a1c3456-6048-4642-9681-aa4d8b7bcc1b.jpg} {Zuna Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/c72897ca-fb6c-43e0-82e9-639db950bf24.jpg} {Dreamsicle Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/4756eda7-7dbd-412b-9c9e-fa3d964321b7.jpg} {Dimple Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/e86c00ce-c6f2-46fd-bb93-d91fb195f8bd.jpg} {Ninja Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/36ddf3b6-fcec-4079-987f-23c180b0e7ca.jpg} {Kinsley Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/fc40398a-6390-4863-86ae-d47772a0b73b.jpg} {Tiny Tina Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/152005a3-46bd-49e3-9395-9e9a441c0156.jpg} {Espresso Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/f6697b78-3c8d-4c31-9845-8642cfd98c2e.jpg} {Meech Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/64522aa3-8bd4-49ad-b5c2-beeb3f0f3185.jpg} {Louise Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/a32a3beb-4ffe-44b7-a1d5-e4496c531bcc.jpg} {Boone Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/af53cad1-491e-45bf-8538-a52c2ad94cbe.jpg} {Brittney Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/d9a25ba3-68c3-49e5-9d54-d7d0906f39d9.jpg} {Billy Joe Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/ad4d086f-1276-481e-bae4-6e127b0467f0.jpg} {Penny Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/1a850656-7b91-46b6-9061-5ee942a13a08.jpg} {Shandor Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/441ed56a-f67a-48b3-b424-34ac8f9a2474.jpg} {Sophia Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/d5502461-9780-42a1-8cc4-f7d5c265bcc4.jpg} {Kolie Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/69bccc10-d941-4e96-9edc-38bd16124067.jpg} {Bindi Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/8123ffeb-2850-4094-9613-26aa019894e0.jpg} {Spritz Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/b875aeea-8638-4d6c-b18d-9b125576a9b8.jpg} {Crackle Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/349a42ab-d1f6-4ba4-9c4c-097f06960545.jpg} {Double Bubble Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/851f6850-b76b-40c3-abca-fbaaa6b22480.jpg} {Maya Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/20242e41-a2b1-4a66-b5a1-b066fa23358b.jpg} {Herman Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/282eb3a6-7077-4a15-ae75-9ec128245252.jpg} {Ron Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/8f005689-f582-4fcf-a755-936739995658.jpg} {Dobby Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/d49170bd-9f82-47d9-b139-76b41127b90b.jpg} {Gobi Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/8490365f-9209-4755-b2a6-6d73e9c3c907.jpg} {Emmanuel Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/85596ac7-c26b-4748-afa6-12cd42d955c3.jpg} {Sonny Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/bed62086-aede-47a6-82b8-59f9e95b5d03.jpg} {Jason Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/dd2239df-420a-4a0d-8100-7226c8b9a7a8.jpg} {Klaus Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/05d835cf-c06a-4466-b363-a73d99379df0.jpg} {Buddha Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/b074bead-119d-418f-bfed-8f69e1cfcaa9.jpg} {Christmas Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/45c72ace-7a80-4b27-a641-d5841ef9a620.jpg} {Wheaties Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/3b9b92ea-01b6-4bf4-b8be-6b9b63e6c378.jpg} {Kimchi Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/4a0cd2af-4af9-4dca-bf5c-4f82e951c7d2.jpg} {Dashi Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/35732d55-54de-4116-a936-8b41fd9e7849.jpg} {Udon Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/d7f792f4-aebc-4846-896c-5d743527293b.jpg} {Falafel Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/905965dc-ae20-4a65-a53c-458673e17888.jpg} {Gelato Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/3e712e32-c2bd-4a8c-8d92-b9e22537a62c.jpg} {Tahini Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/f089a915-8d51-43e6-a62c-365b02d623f6.jpg} {Tortellini Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/0be71903-22b0-42d6-95c1-b6ec72331611.jpg} {Lemon Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/568e854d-9bf1-4ff5-a732-5768b1a8c60c.jpg} {Poppy Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/bb69fec6-1b2f-419f-8b91-56a6a78ae247.jpg} {Kerrigan Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/90bcb7e5-ed8e-4218-a9fc-3a6610d8542f.jpg} {Ed Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/a31e2dcb-ea27-4372-809b-cd39d603fbd8.jpg} {Crepe Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/720736da-d23c-4238-99d9-a70ec664f4be.jpg} {Pita Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/a21ee05c-8a31-4bb7-95cb-4ff93f110bba.jpg} {Nox Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/7c2a89e8-9f3d-42cf-941c-6bc1ee69bddf.jpg} {Nash Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/6292f432-f221-45fb-91c3-af46a7d5e7c8.jpg} {Taos Male Turkish Angora/Mix On Website https://g.petango.com/photos/370/d3a02238-aaf3-48db-89d2-baaef96748a2.jpg} {Mango Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/12ad9a8d-8ac7-4aaa-93d3-516af04ac5e6.jpg} {Tyler Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/f35c3d8d-6dca-400d-8bfe-a33b79e89451.jpg} {Pop Female Domestic Medium Hair/Mix On Website https://g.petango.com/photos/370/e1648f65-0e98-4e21-a645-5d992320b689.jpg} {Rabbit Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/aac355d7-674e-4301-aa46-5f5d8aca8398.jpg} {Velvet Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/80e11cde-a1ae-4c97-a964-8d9d16818593.jpg} {Ornament Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/2b581d6c-97f9-4649-8ecb-a7ce26634972.jpg} {Garland Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/abf5083c-e797-4eb4-b68f-87eb628318d5.jpg} {Candy Star Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/937217b9-bb5d-48b5-9c4d-62f09e23f30d.jpg} {Mello Yello Male Turkish Angora/Mix On Website https://g.petango.com/photos/370/b0cd72ac-7830-4cb6-a814-d6212277655b.jpg} {Simba Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/b9aaef40-00bd-4e6d-9ed5-83a3cc40a8fe.jpg} {Lorelai Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/8bb4c56f-d226-4a3a-a57f-fd3a0c8c230b.jpg} {Zeke Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/ab2cef1d-fc50-4cbc-9237-64151212f268.jpg} {Stetson Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/488cf7a4-0986-426d-a034-53a7b4c8450a.jpg} {Charrito Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/a579aa48-470b-4f1d-bf62-76af285a908f.jpg} {Jaguar Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/3423b3e2-47bd-45c3-819a-5d042e393748.jpg} {Sailboat Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/43b59a54-edc0-45af-b4a9-70117fac50c3.jpg} {Macie Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/6165cdbd-71b9-43a4-a60d-7a1f05e67139.jpg} {Gremlin Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/ea4e2092-d80d-4ec0-9442-44031666dbba.jpg} {Davidson Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/fb39c2e0-e5ab-4a77-8611-c5647a331eaa.jpg} {Mr. Chunks Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/60e3467d-3918-4647-aa67-5e240fcad8ec.jpg} {PeePaw Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/485cad5d-0676-4ea3-b66f-a6509901a80f.jpg} {Phence Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/dede1fe4-bfe1-4652-b798-4257970197fb.jpg} {Garfield Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/d4ac8055-2b49-46d9-b02a-2235ba59bb1e.jpg} {Arlene Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/c58f8b57-489d-4278-b715-cbbec381b6eb.jpg} {Theodore Male Domestic Shorthair/Mix On Website https://g.petango.com/shared/Photo-Not-Available-cat.gif} {Apollo Male Domestic Shorthair/Mix On Website https://g.petango.com/shared/Photo-Not-Available-cat.gif} "
	expected := catsItem
	actual := keyWestString()
	if expected != actual {
		t.Errorf("Expected %s do not match actual %s", expected, actual)
	}
}
func TestInWebsiteMarathon(t *testing.T) {
	catsItem := "{Candi Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/94aae0fb-60b6-40d4-af5e-b1f93049f1b5.jpg} {Emily Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/763df0c2-557c-46a3-a21e-f8b99e29e80b.jpg} {Skyy Female American Shorthair/Mix On Website https://g.petango.com/photos/370/418e2721-4c35-45ef-a61b-c38ca0bb955f.jpg} {Whopper Female Domestic Medium Hair/Mix On Website https://g.petango.com/photos/370/d4d005d1-748e-4fca-9145-d22d3a5dfd69.jpg} {Cornbread Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/5c1e5aca-6a96-4a94-bf69-03db834dedce.jpg} {Sunshine Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/1b9f908e-e095-4e41-9a3a-a38168355058.jpg} {Jethro Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/11329bef-30f3-42db-bdd7-23a3a0a73273.jpg} {Blondi Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/14d2e966-98a0-4771-a0ae-357ee23f25a1.jpg} {Baby Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/25eedb98-0ca6-44d4-8389-d30d8e966f76.jpg} {Cowboy (Golden Paw) Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/5da00e92-1f34-42a1-9ccc-a4997dbc8193.jpg} {Pugsley Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/5ff90e1d-140a-48f2-ae9d-16e39d50fb17.jpg} {Wednesday Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/e694a6f9-ca76-4896-85d2-f1bbf0867999.jpg} {Aria Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/c405ddf6-8297-4904-8be3-cc99e22cca19.jpg} {Doodle Unknown Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/a273d761-b579-4ffe-b7d0-828627b9aeab.jpg} {Cali Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/46a54877-87fb-4535-84a4-4087f286de0d.jpg} {Walnut Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/410bb56e-0cbc-43d9-b8dd-d7b235c9fd38.jpg} {Peanut Female Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/654d9a90-309c-424e-86d3-4212b4eef991.jpg} {Egan Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/99a8f8be-f306-4ede-b289-04c2831dbd65.jpg} {Keagan Male Domestic Shorthair/Mix On Website https://g.petango.com/photos/370/bf1b1beb-8152-4110-9925-2c424fdf4cc8.jpg} {On Website On Website On Website On Website On Website} {On Website On Website On Website On Website On Website} "
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

// --------------------------------------------------------Sprint 3 Tests----------------------------------------------------------------
func TestFeaturesMiamiDade(t *testing.T) {
	var catList []Cat

	m := map[string]int{
		"friendly":        1,
		"playful":         2,
		"adorable":        3,
		"loving":          4,
		"affectionate":    5,
		"curious":         6,
		"energetic":       7,
		"loves attention": 8,
		"loves to cuddle": 9,
		"loves to play":   10,
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
func TestFeaturesLakeCounty(t *testing.T) {
	var catList []Cat

	m := map[string]int{
		"friendly":        1,
		"playful":         2,
		"adorable":        3,
		"loving":          4,
		"affectionate":    5,
		"curious":         6,
		"energetic":       7,
		"loves attention": 8,
		"loves to cuddle": 9,
		"loves to play":   10,
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
func TestFeaturesPeggy(t *testing.T) {
	var catList []Cat

	m := map[string]int{
		"friendly":        1,
		"playful":         2,
		"adorable":        3,
		"loving":          4,
		"affectionate":    5,
		"curious":         6,
		"energetic":       7,
		"loves attention": 8,
		"loves to cuddle": 9,
		"loves to play":   10,
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
