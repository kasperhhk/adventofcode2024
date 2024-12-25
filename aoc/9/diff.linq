<Query Kind="Statements">
  <Namespace>System.Threading.Tasks</Namespace>
</Query>

var p = @"C:\Users\kaspe\Documents\dev\GIT\adventofcode2024\aoc\8\";
var q1 = File.ReadAllLines(p + "Q1.txt");
var q2 = File.ReadAllLines(p + "Q2.txt");

if (q1.Length != q2.Length)
	throw new Exception("len ");
	
q1.Length.Dump();
	
for (var i = 0; i<q1.Length; i++) {
	if (q1[i] != q2[i]) {
		"------------------".Dump();
		$"i={i}".Dump();
		$"CORRECT={q1[i]}".Dump();
		$"WRONG={q2[i]}".Dump();
	}		
}