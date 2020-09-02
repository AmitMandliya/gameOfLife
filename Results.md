### Methods

To identify the easier of all the languages (Go, Kotlin and Julia), we conducted face to face interviews with 10 participants. We coded Game of life in all the three languages and introduced multiple bugs in each language. These were the logical bugs so that a participant who is new to the language does not have to search for syntax in a limited time period of interviews. <br>
Our study with the participants in a limited time frame needed us to get everything related to set up done quickly. We decided to use the feature of repl.it where the participants just have to click a link and the buggy code will be ready for with a terminal on the right hand side of the screen. <br>
We asked the participants to share their screen while they find the bugs and noted down below stuff for each one of them:
- Number of time the code was executed <br>
- Number of time it took for them to find the first bug <br>
- Number of google searches they had to make. <br>
At the end of each interview, we asked the participants to fill out a Google Doc form which asked them about the familiarity of the language and certain other questions described in further sections. <br>
After all the interviews were over, we did an analysis of the results to come up with the best of all the three languages. <br>


### Materials

1. Google Form

The Google form was created to record testers’ responses. Each tester was asked to fill it up after the completion of their debugging session. The form questioned the testers about the following:
-Level of familiarity with individual languages
-Level of difficulty in finding the bugs
-Their approach towards debugging for each language they chose
-How obvious were the bugs to them?

The responses to these questions helped us better analyze the general obscurity of the languages, and also understand how hard was it to find the bugs that we introduced.

2. Google Sheet

The Google sheet was created to record certain performance metrics of testers during their debugging sessions. We recorded the following metrics:
-The Time(minutes) taken by tester to debug the first bug: It gave us an idea about how much time does it take for a tester to get familiar enough with a language to debug a bug.
-Number of execution runs of the code: This metric helped us infer how hard it was to spot all the bugs in the code for the tester.
-Number of Google Searches: This metric helped us get a clearer idea about tester’s familiarity with the language.


### Observations

### Inferences

### Conclusions

### Threats to validity
- Our experiment involved speaking with people for 30 minutes each and asking them to find bugs in codes written in languages they have probably never worked in before. We had 3 such codes and realized that it is difficult for the testers to be able to complete debugging the three codes in merely 30 minutes. While some of the testers were able to do it, most were not. Because of this the data collected is insufficient. <br>
- Some of the participants, for the previous homework, had worked on the same code with the same logic in the same language. So they were able to find the syntactical and logical errors a little more easily despite barely knowing the coding language. Had we been working on different games with different logics, they may have taken more time to find the bugs. <br>
- The commit history was public, so that could have been referred to before or during the debugging sessions. <br>
