package handler

const IndexPage = `
<html>
	<head>
		<title>OAuth-2 Test</title>
	</head>
	<body>
		<h2>OAuth-2 Test</h2>
		<p>
			Login with the following,
		</p>
		<ul>
			<li><a href="/login-gl">Google</a></li>
			<li><a href="/login-fb">Facebook</a></li>
		</ul>
	</body>
</html>
`

const TokenPage = `
<!DOCTYPE html>
<html>
<body>
<a href="/">Home</a>
<h3>Token Generated</h3>
<h5>Token:<h5>
<textarea  rows="4" cols="50" id="myInput" disabled>{{ .AccessToken }}</textarea>
<button onclick="myFunction()">Copy</button>
<span id="copied"></span>
<script>
function myFunction() {
  var copyText = document.getElementById("myInput");
  copyText.select();
  copyText.setSelectionRange(0, 99999);
  navigator.clipboard.writeText(copyText.value);
  var s = document.getElementById("copied");
  s.innerHTML = "Copied!!"
}
</script>
</body>
</html>
`
