function openNav() {
    document.getElementById("mySidenav").style.width = "250px";
  }
  
  function closeNav() {
    document.getElementById("mySidenav").style.width = "0";
  }
  function ValidateEmail(inputText)
  {
  var mailformat = /^w+([.-]?w+)*@w+([.-]?w+)*(.w{2,3})+$/;
  if(inputText.value.match(mailformat))
  {
  document.form1.text1.focus();
  return true;
  }
  else
  {
  alert("Entered Email address is invalid!");
  document.form1.text1.focus();
  return false;
  }
  }

  