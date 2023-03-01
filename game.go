package main

import (
	"log"

	"example.com/m/elements"
	"github.com/gdamore/tcell/v2"
)

func getElements(px, py int, cameraCallback func(int, int)) []elements.Element {
	wm, err := ReadMap("map.txt", map[rune]elements.Tileable{
		'T': elements.NewSimpleTile(" ^ |\n/ \\ \n/ \\ \n | ^\n\\ / \n\\ / ", tcell.StyleDefault.Foreground(tcell.ColorLightGreen)),
	})

	if err != nil {
		panic(err)
	}

	txt := elements.NewTextbox()

	p := elements.NewPlayer(px, py, cameraCallback)

	log0 := elements.NewSign(38, 25, []string{"HAWK PEAK TRAIL",
		"Since 2023 started, I've been keeping a google sheet to track how I spend my time.",
		"It's definitely changed how I've both thought of my life, as well as how I've conducted it, and shown me the habits I didn't know I had.",
	}, txt, p)
	log1 := elements.NewSign(34, 33, []string{"A DAY/WEEK/MONTH IN THE LIFE",
		"I've realized my life is mostly projects and youtube, which is both interesting and depressing.",
		"So far, the projects have been 1. extending an open source project (the engine I created this in!) to work on the web,",
		"2. conlanging (which didn't go anywhere), 3. robotics (the most pervasive and time consuming), 4. summer program applications (just kind of takes time),",
		"5. bannai snake, a small game I've been wanting to make, and 6. some hardware development I'm currently working on.",
		"I think those are all valid, but my time between projects always sinks into pure youtube days, where it's literally sleep, food, and videos.",
		"I want to wean off of youtube in the future, but the only way to do that right now seems to be other projects",
		"and I want a better balance between recreation and working on stuff going into the future.",
	}, txt, p)
	log2 := elements.NewSign(47, 39, []string{"DERIVATIVE OF THOUGHT OVER TIME",
		"For better or for worse, I now conduct my life in 30 minute chunks- I now make sure I end an activity when the time ends in XX:00 or XX:30, because that's easier to track in my log.",
		"It sometimes makes me more productive, like when I start programming and then \"lock myself\" into doing it for another ten minutes",
		"but it sometimes also \"forces\" me to watch youtube for another fifteen after I finish a video.",
	}, txt, p)
	log3 := elements.NewSign(44, 46, []string{"FIFTEEN MINUTE SYNDROME",
		"Logging in 30 minute increments causes me to pivot to productive things when the time hits XX:15 and XX:45",
		"because I don't want to log the unproductive things I do. This has led to some of the most productive days of my life.",
		"But some days, I give up. I relegate myself to watching youtube for four to six hours a day, especially between projects,",
		"and I feel like my days have been swinging wildly between productive and unproductive since I started the schedule.",
	}, txt, p)

	yoji0 := elements.NewSign(27, 16, []string{"IT'S NOT ALWAYS BLACK AND WHITE (but in this case it is)",
		"In my Japanese class, I watched Yojimbo, a 1961 samurai film that inspired many spaghetti westerns, and directed by Akira Kurosawa.",
		"I think it was the first full length black and white film I've watched, and I was really surprised at how much I enjoyed it. I thought it would be really hard to tell what was going on, but it was executed very well and had a interesting story.",
		"It also had a lot of rhetorical devices, which I'll be mostly focusing on here.",
	}, txt, p)
	yoji1 := elements.NewSign(17, 13, []string{"THE ORIGINAL SPAGHETTI WESTERN VILLAIN",
		"Unosuke, a gang leader's brother, is a gunslinging, dashing, and deceptive character. He is often conflated with the wind, and is almost always depicted with his trusty pistol.",
		"Whenever Unosuke arrives, he is usually preceded by wind, which I think indicates that he is fickle- not really grounded to ideals, but going with the flow and using charisma and guile to achieve his goals.",
		"The protagonist (who remains unnamed but goes by \"mulberry 30 year old\"), when preparing to confront Unosuke, literally throws a knife at a leaf in the wind, further reinforcing this connection.",
		"Unosuke's pistol also is with him at most points, where he reveals it from inside his shirt, signaling that violence is an integral part to his identity, like another limb- a part of his body.",
		"He also begs for it when dying and says he \"feels naked without his gun\", further showing how violence is part of his character, and then proceeds to attempt to shoot the protagonist before dying anyways. Unosuke died how he lived, trying to inflict violence on others.",
	}, txt, p)
	yoji2 := elements.NewSign(21, 6, []string{"CLINT EASTWOOD'S INSPIRATION",
		"The main protagonist, who remains unnamed, has a very strange relationship with honor. He never expects it in others, and even encourages cowardice sometimes, but also has a strong moral code for himself.",
		"For example, when one of the hired bodyguards decides to ditch after he gets paid, the protagonist encourages him to leave. This is out of his own self-interest (he also plans on leaving when the battle starts because he overheard the gang trying to kill him later)",
		"but I think it goes beyond that. The protagonist's stated goal is to make the town peaceful again, so he's generally opposed to violence, but commits it himself anyways as kind of a preventative measure for others. He therefore is usually all for cowardice in others when it avoids violence, but sharply critical when that cowardice makes others suffer.",
		"Another instance of this is when he calls the gang leaders cowards at certain points throughout the movie. Their cowardice leads to many of the violent acts seen throughout the movie and the suffering of the townsfolk, and in this case, he berates them for their lack of honor.",
	}, txt, p)

	bye0 := elements.NewSign(49, 19, []string{"THE ONLY GAME THAT MADE ME CRY",
		"I don't really cry. I often ask people how often they cry: some people say every other day, some say once a week, some say once in a few months. I cry every two-ish years. It's just not something I do a lot, and especially not at trivial things like books, movies, or games.",
		"I have watched a playthrough of BEFORE YOUR EYES three times before. By the time I get to the end, I am always silently crying my heart out. It's a feeling I never really get in normal life, and sometimes I want to watch it again just to feel that gut wrenching emotion again. I'm not exactly sure why.",
		"I took notes mostly on the rhetorical analysis side of the game, but I thought I might focus on the themes and how they've impacted my outlook on life, with a little rhetorical analysis on the side for this activity.",
	}, txt, p)
	bye1 := elements.NewSign(71, 20, []string{"THE PASSING DAY",
		"I'm going to try and not spoil the story at all, but I'm not sure how that'll go, so please play/watch the game first. There's this one moment at the end of the game where Benny, the main character, is handed a piece of music and told that playing it when he was four helped his parents through a tough time in their lives.",
		"The themes of the story center around legacy, and I'm not sure exactly what causes me to start crying there, but it does. Music throughout the game serves as the passion that Benny's parents have and what path they put him on, while art symbolizes a passion he grows into.",
		"Benny doesn't become a professional musician, like his parents wanted him to. But his actions, even if not exceptional to the rest of the world, WAS exceptional to his parents and made their lives just a bit better. His impact was local, and he won't be remembered as a great musician for the rest of time",
		"but he will have made an impact on those around him, and that's what makes a great life for anybody.",
	}, txt, p)
	bye2 := elements.NewSign(90, 17, []string{"A CHAIN IMPACT",
		"I've always struggled with what I wanted my impact to be on the world. I know it's cliche, but I go into these nihilistic ruts where I ask myself what am I going to be remembered for a century, a decade, a few years after I die. I am almost certainly not going to be remembered like Albert Einstein or Amelia Earheart or Stalin or Steve Jobs. And to be honest, I don't know anything my great-grandmother did or even her name.",
		"I often think about careers that would lead to me being known (artist, CEO, influencer, politician, etc.), but none of them really spark joy in me. I don't think I want to be known during my lifetime; I enjoy my privacy, and it just seems like a pain to constantly be thinking about my image to the outside world.",
		"But what's the point otherwise? That's the question I pondered for a while until I came to the conclusion that I need to be part of a project that impacts humanity. I don't have to be in the spotlight, but if I can be in the end credits of a groundbreaking achievement, that's enough for me, I thought.",
		"The first time I watched Before Your Eyes was a month-ish after my dog passed away. I initially related the story to him, and his legacy on earth. He wasn't famous- he didn't have an instagram account, probably less than a hundred people had seen him throughout his life.",
		"But he had a significant impact on my life: I sulked next to him whenever I felt down, I buried my face in his stomach everyday after school, I walked, I slept, I walked with him. That legacy shapes who I am, and therefore shapes everybody that I have an impact on, passing on that legacy.",
		"Before Your Eyes helped me realize that a life without doing something you'll be remembered for by the world isn't a wasted life. Malt was a part of my life, and his life wasn't irrelevant; as long as I make an (positive) impact on those around me, I don't think my life will be irrelevant, and I think the game helped me solidify my worldview on that.",
	}, txt, p)

	jour0 := elements.NewSign(41, 14, []string{"NO PHONES ALLOWED",
		"Well, I did some journaling! The last time I did something like it was when I was like nine, so I thought I'd give it another go. New rules though- 35 minutes, no phone. I guess I didn't have a phone in third grade, so old rules?",
		"I'm just going to go over the main points I touched on during the writing, as well as my takeaways and inferences from what I talked about. I have a doc with pictures of whole thing (as well as notes) too, just ask if you want to see it!",
	}, txt, p)
	jour1 := elements.NewSign(51, 11, []string{"HIGH SCHOOL DOES MATTER?",
		"Since I wrote the journal a few days after I met up with friends, I was thinking a lot about my (perhaps lacking) social life. I keep on hearing about stories of my parents doing things in their childhood and I worry that I'm not really doing any of those things and when I grow up, my childhood will be summed up as \"got good grades\".",
		"I kind of disagree with Mrs. Salvador's saying that everything you do in high school is irrelevant. My parents have stories of traveling across India, crossing rope bridges while going to school, getting absolutely destroyed for sleeping in class because the teacher's voice sounded like a lullaby. I don't have any stories like that; I just kind of did well in school, played video games, watched youtube, programmed. Not very story worthy.",
		"So I do think the things you do in high school are mostly irrelevant once you get out of high school externally, but internally it builds memories and thereby personalities. I worry I don't have any of those core memories that shaped my parents, and no lifelong bonds that people tend to form with schoolfriends.",
		"And I think part of that is I live in a very sheltered environment. This goes back to that adversity essay, but I wouldn't trade my privilege for \"building character\" or anything like that. But I wonder if I'm missing out on gaining, say, a work ethic because I haven't had to study for tests in years to do ok in a class, or a passion by not focusing on one thing, like percussion. Nothing has forced me to make risks (which lead to memorable moments), so I didn't, and now I have a boring childhood.",
		"It's probably good I have had a boring childhood. \"May you live in interesting times\" is supposed to be a wish of ill will, because interesting times are rife with adversity. You can't write a story without conflict, but I don't want to have conflict in my life, so I guess I'll try not to be in a story, and my childhood was definitely not a story. Hopefully it won't be in the future? I don't really know anymore.",
	}, txt, p)
	jour2 := elements.NewSign(52, 5, []string{"TYPICAL SOCIAL ANXIETY STUFF",
		"Ok, the title might be a bit of a hyperbole. I enjoy having people over in some situations, but I really struggle with initiating hangouts. I'm always... reluctant to ask people to go do something, because I'm afraid I'm too boring to talk to. Before covid, I was definitely an outgoing guy, but afterwards I kind of forgot how to talk to people.",
		"One day during quarantine, for some bizarre reason, I decided to try and see little I could talk and still get by. It was during covid, so it was just my parents and sister in the house, so I'd just kind of be in my room the entire day except for meals. I don't know why I played that stupid game at all, and it really impacted my ability to hold a conversation. It was... entertaining? I guess? The boredom of quarantine?",
		"It's also a general apathy for going out too, because it had been a while since I went out with friends. I actually just went walking over break again with my dog and two of my friends, and also watched a movie with like five other people and also walked to Koja Kitchen.",
		"It was really fun, but before break I kind of forgot how fun it was to be with my friends. I think watching netflix cereals where people hang out, as weird as that sounds, reinvigorated my desire to hangout.",
	}, txt, p)

	ted0 := elements.NewSign(55, 33, []string{"THE ANTI TED TALK TED TALK",
		"Most of my other activities for this project have been related to my outlook on life and what I'm doing with my time. So to finish this project off, I wanted to see how I could actually start implementing the things I've learnt and want to do.",
		"I watched \"Enter the cult of extreme productivity\" from Mark Adams, expecting some motivational talk to get me to try a new way of organizing my life for a day before returning to the lazy self I was before. It turned out to be a kind of self-defeating, yet sort of helpful \"anti-ted talk\", if that makes sense.",
	}, txt, p)
	ted1 := elements.NewSign(69, 30, []string{"DERIVATIVE OF PRODUCTIVITY OVER TIME = ZERO",
		"The main point of the Ted Talk was basically that most self-help books/videos/ted talks advocate for change in thinking over time- that we strive towards a \"perfect version\" of ourselves, reminiscent of Plato's Theory of Forms. But this ideal never comes, because, forgive my math analogy, the derivative of change at today is zero; the derivative in two hours will be zero; the derivative a year from now is zero.",
		"By striving for that ideal, productive version of oneself, you give yourself the excuse of not doing it now, and pushing that change to later, to tomorrow, which never comes. He instead advocates for a non-continuous approach, what he calls \"lock-in\"- changing your behavior immediately, then going from there, in a stair shaped way- because the productive you is not related to the non-productive you.",
		"I don't think one part of his talk is really achievable. What I got was he advocates changing to a critical-thinking self immediately from your experiential self, which isn't as easy as he says it is. I think maybe taking small, IMMEDIATE steps to make the activation energy for doing something unproductive higher might be more achievable, therefore making many small steps, where the derivative at any point is zero or undefined.",
	}, txt, p)
	ted2 := elements.NewSign(86, 31, []string{"FEAR OF FAILURE?",
		"Adams also speaks about the fear of failure, and how contracts fare much better than new year's resolutions. Basically, he advocates for systems that make you fear unproductiveness (or the consequences you put in place for it) and therefore make you strive for productiveness.",
		"He uses Cortez burning his ships as an analogy for this fear, where his men couldn't go back to Spain because there literally wasn't a way to go back. Cortez upped the stakes and made it so they didn't have an escape plan- they had to succeed in conquistadoring, or else they'd die. A similar thing might be enforcing punishments on yourself (like not watching something) if you don't meet a goal.",
		"I'm not sure I wholeheartedly agree with this either. I feel like the negative reinforcement might not work as well as just making it really hard to do some task, which in a sense is punishing you though a time penalty.",
		"For example, I decided to install a site blocker on my laptop; Youtube is now blocked. It's not going to stop me from going there, but while I'm going through the motions of unblocking youtube, I'm hoping that I think about what I'm doing and allow my critical self to realize that I should be doing something more productive.",
	}, txt, p)

	return []elements.Element{
		wm,
		log0, log1, log2, log3,
		yoji0, yoji1, yoji2,
		bye0, bye1, bye2,
		jour0, jour1, jour2,
		ted0, ted1, ted2,
		p, txt}
}

func initializeScreen() tcell.Screen {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.Clear()
	s.SetSize(80, 24)

	return s
}
