## Methods

To identify the simplicity of all the languages (Go, Kotlin and Julia), we conducted face to face interviews with 10 participants. We coded Game of life in all the three languages and introduced multiple bugs in each language. These were the logical bugs so that a participant who is new to the language does not have to search for syntax in a limited time period of interviews (30 min). <br>
We split ourselves into groups of two and conducted the debugging session for each of the testers. One person from the group was actively involved with the tester as he navigated through the code and the other person at the same time noted the performance metrics in a Google Sheet. This helped in critical measurement of the activity. <br> 
Our study with the participants in a limited time frame needed us to get everything related to set up done quickly. We decided to use the feature of repl.it where the participants just have to click a link and the buggy code will be ready for with a terminal on the right hand side of the screen. <br>
We asked the participants to share their screen while they find the bugs and noted down below stuff for each one of them:
- Number of time the code was executed <br>
- Number of time it took for them to find the first bug <br>
- Number of google searches they had to make. <br>

At the end of each interview, we asked the participants to fill out a Google Doc form which asked them about the familiarity of the language and certain other questions described in further sections. <br>  

After all the interviews were over, we did an analysis of the results to come up with the best of all the three languages. <br>


## Materials

1. Google Form

The Google form was created to record testers’ responses. Each tester was asked to fill it up after the completion of their debugging session. The form questioned the testers about the following:
- Level of familiarity with individual languages
- Level of difficulty in finding the bugs
- Their approach towards debugging for each language they chose
- How obvious were the bugs to them?

The responses to these questions helped us better analyze the general obscurity of the languages, and also understand how hard was it to find the bugs that we introduced.

2. Google Sheet

The Google sheet was created to record certain performance metrics of testers during their debugging sessions. We recorded the following metrics:  
- The Time(minutes) taken by tester to debug the first bug: It gave us an idea about how much time does it take for a tester to get familiar enough with a language to debug a bug.
- Number of execution runs of the code: This metric helped us infer how hard it was to spot all the bugs in the code for the tester.
- Number of Google Searches: This metric helped us get a clearer idea about tester’s familiarity with the language.


## Observations

Based on the testers' survey reponses and first hand data collected, we have come to following inferences
- Julia came across as the least preferred language whereas Kotlin was the preferred choice of language for Game Of Life
- Although more than half of the testers were not familiar to the language, average time to find the first bug was least in Julia (5.71s) followed by Go.
- Kotlin required more time to find the first bug but it had least amount of executions that the participants did to resolve all the bugs.
- The bugs that were introduced in the code weren't obvious and the the participants rated it moderate on level of difficulty
- Many of the participants referred to life.awk as a reference for debugging the code
- Not understanding syntax was one of the major hurdles that were faced by them

**[Data Analysis done can be found here](Analysis_Of_Responses.ipynb)**. 

### Conclusions
The testers' responses as well as our observations during each of the testing session have brought us to a conclusion that two person per participant is the most effective way for quality measurement of the activity and proper planning of the activity is crucial for acurate measurements. 
From the point of view of implementing the Game of life in uncommon languages, we have come to a conclusion that Julia is easier to debug than Go and Kotlin.


### Threats to validity
- Our experiment involved speaking with people for 30 minutes each and asking them to find bugs in codes written in languages they have probably never worked in before. We had 3 such codes and realized that it is difficult for the testers to be able to complete debugging the three codes in merely 30 minutes. While some of the testers were able to do it, most were not. Because of this the data collected is insufficient. <br>
- Some of the participants, for the previous homework, had worked on the same code with the same logic in the same language. So they were able to find the syntactical and logical errors a little more easily despite barely knowing the coding language. Had we been working on different games with different logics, they may have taken more time to find the bugs. <br>
- The commit history was public, so that could have been referred to before or during the debugging sessions. <br>
- Many of the participants referred to life.awk file as a reference for resolving the bug which we feel do not correctly resemble the testers' debugging skills.  Since, the same code were re-written in different languages, find bugs was easier for them. Had they been just knowing the algorithm, the debugging would have been a accurate reflection of the skills.
