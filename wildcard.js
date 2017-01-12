//github.com/ZoneTong/wildcard
// wildcard like in Makefile
// simple support rules: * ? [...]
// not support: '\'
function ExactMatch(pattern, str){
	var lp = pattern.length;
	var ls = str.length;
	if (lp == 0 && ls == 0){
		return true;
	} else if (lp == 0 || ls == 0){
		return false;
	}
	
	var p = pattern[0];
	if (p == '*') {
		for (var i = 0; i <= ls; ++i){
			if (ExactMatch(pattern.substr(1), str.substr(i))){
				return true;
			}
		}
	} else if (p == '?'){
		return ExactMatch(pattern.substr(1), str.substr(1));
	} else if (p == '['){
		var i = pattern.indexOf(']');
		if (i >= 0 && pattern.substring(1, i).indexOf(str[0]) >= 0){
			return ExactMatch(pattern.substr(i+1), str.substr(1));
		}
	} else if (p == str[0]){
		return ExactMatch(pattern.substr(1), str.substr(1));
	}
	return false;
}