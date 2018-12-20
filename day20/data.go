package main

func getInput() string {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(content)
}

var inputStr string = getInput()

const (
	testStr1 string = `^WNE$`
	testStr2 string = `^ENWWW(NEEE|SSE(EE|N))$`
	testStr3 string = `^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$`
	testStr4 string = `^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$`
	testStr5 string = `^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$`
)
