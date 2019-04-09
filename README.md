# Go
Go lang notes and Programs

Problem statment: 

Shlok and Sachin are good friends. Shlok wanted to test Sachin, so he wrote down a string S with length N and one character X. He wants Sachin to find the number of different substrings of S which contain the character X at least once. Sachin is busy with his girlfriend, so he needs you to find the answer.

Two substrings of S are considered different if their positions in S are different.

Input
The first line of the input contains a single integer T denoting the number of test cases. The description of T test cases follows.
The first line of each test case contains a single integer N.
The second line contains a string S with length N, followed by a space and a character X.
Output
For each test case, print a single line containing one integer — the number of substrings of S that contain X.

Constraints
1≤T≤1,000
1≤N≤106
S contains only lowercase English letters
X is a lowercase English letter
the sum of N over all test cases does not exceed 106
Subtasks
Subtask #1 (30 points): the sum of N over all test cases does not exceed 103
Subtask #2 (70 points): original constraints

Example Input
2
3
abb b
6
abcabc c
Example Output
5
15
Explanation
Example case 1: The string "abb" has six substrings: "a", "b", "b", "ab", "bb", "abb". The substrings that contain 'b' are "b", "b", "ab", "bb", "abb".
