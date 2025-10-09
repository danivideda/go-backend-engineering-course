package db

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/danivideda/go-backend-engineering-course/social/internal/store"
)

var usernames = []string{
	"bob", "alice", "dani", "jared", "emma", "liam", "olivia", "noah", "sophia", "ethan", "ava", "mason", "isabella",
	"james", "mia", "benjamin", "charlotte", "henry", "amelia", "alexander", "harper", "michael", "evelyn", "daniel", "abigail",
	"jacob", "emily", "william", "elizabeth", "david", "sofia", "joseph", "avery", "thomas", "ella", "christopher", "scarlett",
	"matthew", "grace", "andrew", "chloe", "samuel", "victoria", "gabriel", "lily", "nicholas", "madison", "john", "zoey", "steven",
}

var emails = []string{
	"bob.smith", "alice.jones", "dani.brown", "jared.lee", "emma.wilson", "liam.taylor", "olivia.moore", "noah.jackson", "sophia.martin",
	"ethan.white", "ava.thomas", "mason.miller", "isabella.davis", "james.roberts", "mia.jackson", "benjamin.clark", "charlotte.lewis", "henry.walker",
	"amelia.hall", "alexander.allen", "harper.young", "michael.hernandez", "evelyn.king", "daniel.wright", "abigail.scott", "jacob.green", "emily.baker", "william.gonzalez",
	"elizabeth.nelson", "david.carter", "sofia.mitchell", "joseph.perez", "avery.roberts", "thomas.turner", "ella.phillips", "christopher.campbell", "scarlett.parker", "matthew.evans",
	"grace.edwards", "andrew.collins", "chloe.stewart", "samuel.morris", "victoria.rogers", "gabriel.reed", "lily.cook", "nicholas.morgan", "madison.bell", "john.murphy",
	"zoey.rivera", "steven.cooper",
}

type postMock struct {
	Title   string
	Content string
	Tags    []string
}

var postsMock = []postMock{
	{Title: "Advancements in Technology", Content: "Technology is advancing at an unprecedented pace, reshaping industries and daily life. From AI breakthroughs to quantum computing, the potential for innovation is limitless, but so are the challenges of accessibility and equity.\n\nAdapting to these changes requires continuous learning and ethical considerations. How can we ensure technology benefits all? Engaging in open discussions about its societal impact is a great starting point.", Tags: []string{"Technology", "Innovation", "Ethics"}},
	{Title: "Creativity at the Intersection of Art and Technology", Content: "Creativity thrives at the intersection of art and technology. Digital tools like graphic tablets and 3D modeling software enable artists to push boundaries, creating immersive experiences that were once unimaginable.\n\nExperimenting with these tools fosters unique styles and new forms of expression. Aspiring creators should explore platforms like Procreate or Blender to unlock their potential and share their work with global communities.", Tags: []string{"Art", "Digital Tools", "Creativity"}},
	{Title: "The Benefits of Nature Walks", Content: "Nature walks offer a profound way to reconnect with the environment. Observing wildlife and seasonal changes on trails provides a refreshing escape and a reminder of our planet's beauty and fragility.\n\nConservation efforts are vital to preserving these natural spaces for future generations. Supporting local initiatives or reducing personal waste can make a tangible difference in protecting ecosystems.", Tags: []string{"Nature", "Conservation", "Hiking"}},
	{Title: "The Role of Sports in Personal Growth", Content: "Sports foster resilience, teamwork, and physical health. Whether playing soccer or running marathons, the discipline learned on the field translates to personal and professional growth.\n\nEngaging in regular physical activity boosts mental clarity and builds community. Joining local leagues or casual groups can enhance both fitness and social connections.", Tags: []string{"Sports", "Teamwork", "Fitness"}},
	{Title: "The Power of Reading for Personal Development", Content: "Reading books opens doors to new perspectives and knowledge. From historical novels to scientific journals, each page deepens understanding and empathy for diverse experiences.\n\nSharing book recommendations with friends sparks meaningful conversations. Building a habit of reading diverse genres can enrich personal growth and inspire lifelong learning.", Tags: []string{"Reading", "Literature", "Empathy"}},
	{Title: "How Music Shapes Our Emotions", Content: "Music shapes emotions and experiences, from soothing classical to upbeat pop. Creating playlists for work, relaxation, or celebrations enhances the moment and connects us to memories.\n\nLive concerts unite people in shared joy, while producing your own tracks can be a creative outlet. Exploring genres or learning an instrument opens new avenues for expression.", Tags: []string{"Music", "Creativity", "Concerts"}},
	{Title: "Exploring Global Cuisines Through Cooking", Content: "Cooking with fresh ingredients transforms meals into cultural experiences. Exploring cuisines from Italian to Thai introduces new flavors and techniques that elevate home dining.\n\nHosting dinner parties with global recipes fosters connection and storytelling. Experimenting with spices or local markets can turn every meal into an adventure.", Tags: []string{"Cooking", "Cuisine", "Community"}},
	{Title: "Mastering the Art of Photography", Content: "Photography captures moments that tell stories through light and composition. Whether shooting landscapes or portraits, the right angle can convey profound emotions or narratives.\n\nEditing software enhances images, adding mood or clarity. Sharing your work online or at exhibits inspires others to see the world through your lens.", Tags: []string{"Photography", "Art", "Editing"}},
	{Title: "Yoga for Balance and Wellness", Content: "Yoga promotes physical and mental balance through intentional movement and breath. Regular practice enhances flexibility, reduces stress, and fosters a sense of calm.\n\nGroup classes or online sessions build community and accountability. Starting with simple poses like downward dog can lead to a lifelong journey of wellness.", Tags: []string{"Yoga", "Wellness", "Mindfulness"}},
	{Title: "The World of Gaming and Esports", Content: "Gaming offers immersive adventures and strategic challenges. From open-world RPGs to competitive esports, players develop quick thinking and collaboration skills.\n\nMultiplayer platforms connect global communities, fostering friendships. Exploring game design or streaming can turn a hobby into a creative career.", Tags: []string{"Gaming", "Esports", "Community"}},
	{Title: "The Joys of Home Gardening", Content: "Gardening nurtures both plants and patience, creating vibrant spaces even in small yards. Growing your own herbs or vegetables connects you to the cycles of nature.\n\nSustainable practices like composting reduce waste and enrich soil. Sharing harvests with neighbors builds community and promotes eco-friendly living.", Tags: []string{"Gardening", "Sustainability", "Nature"}},
	{Title: "The Transformative Power of Travel", Content: "Travel exposes us to diverse cultures and breathtaking landscapes. Planning trips, whether local or international, sparks excitement and broadens worldviews.\n\nImmersing in local customs, like trying street food or learning phrases, creates lasting memories. Documenting journeys through blogs or photos preserves the experience.", Tags: []string{"Travel", "Culture", "Adventure"}},
	{Title: "Expressing Emotions Through Painting", Content: "Painting transforms emotions into visual stories. Whether using oils or watercolors, the act of creating art fosters self-expression and mindfulness.\n\nExhibiting work in galleries or online platforms connects artists with audiences. Drawing inspiration from nature or personal experiences fuels unique creations.", Tags: []string{"Painting", "Art", "Expression"}},
	{Title: "Coding as a Tool for Problem-Solving", Content: "Coding empowers problem-solving through logic and creativity. Building apps or automating tasks demonstrates the practical impact of programming in daily life.\n\nContributing to open-source projects or learning frameworks like Go or Python accelerates growth. The tech community thrives on shared knowledge and innovation.", Tags: []string{"Coding", "Programming", "Open Source"}},
	{Title: "Dancing as a Universal Language", Content: "Dancing celebrates movement and rhythm, from salsa to contemporary styles. It’s a universal language that expresses joy and builds physical strength.\n\nJoining dance classes or performing in recitals creates bonds and confidence. Exploring choreography or freestyle opens up endless creative possibilities.", Tags: []string{"Dance", "Movement", "Creativity"}},
	{Title: "Lessons from History for Today", Content: "History offers lessons from past events that shape our present. Studying eras like the Renaissance or modern revolutions reveals patterns in human behavior.\n\nVisiting museums or reading primary sources deepens understanding. Applying historical insights helps navigate today’s societal challenges thoughtfully.", Tags: []string{"History", "Learning", "Society"}},
	{Title: "The Impact of Volunteering", Content: "Volunteering creates positive change in communities while fostering personal growth. Whether helping at shelters or organizing cleanups, every effort counts.\n\nCollaborating with others on meaningful projects builds lasting connections. Finding causes you’re passionate about amplifies your impact.", Tags: []string{"Volunteering", "Community", "Impact"}},
	{Title: "Exploring the Wonders of Astronomy", Content: "Astronomy unveils the universe’s wonders, from distant galaxies to nearby planets. Stargazing with a telescope or app sparks curiosity and awe.\n\nFollowing space missions or joining astronomy clubs deepens engagement. The cosmos invites us to explore questions about our place in it.", Tags: []string{"Astronomy", "Space", "Exploration"}},
	{Title: "The Craft of Writing and Storytelling", Content: "Writing captures thoughts and stories, from journals to novels. Putting words to paper clarifies ideas and preserves personal or fictional narratives.\n\nSharing drafts with writing groups or publishing online invites feedback. Consistent practice hones a unique voice and storytelling craft.", Tags: []string{"Writing", "Storytelling", "Creativity"}},
	{Title: "Building Fitness Through Consistent Routines", Content: "Fitness routines enhance strength and mental clarity through discipline. Mixing cardio, weights, or bodyweight exercises keeps workouts engaging and effective.\n\nPairing exercise with balanced nutrition maximizes results. Joining fitness challenges or apps can sustain motivation and track progress.", Tags: []string{"Fitness", "Health", "Exercise"}},
	{Title: "Sustainable Fashion and Personal Style", Content: "Fashion reflects individuality through evolving trends and personal style. Choosing sustainable brands or thrifting reduces environmental impact while staying chic.\n\nExperimenting with accessories or bold colors creates signature looks. Fashion blogs or shows inspire creative outfit ideas.", Tags: []string{"Fashion", "Sustainability", "Style"}},
	{Title: "The Art of Filmmaking and Storytelling", Content: "Films tell powerful stories through visuals and sound. From indie dramas to blockbusters, directors craft narratives that provoke thought or entertain.\n\nAttending film festivals or analyzing classics deepens appreciation. Sharing reviews online sparks discussions about cinematic art.", Tags: []string{"Film", "Cinema", "Storytelling"}},
	{Title: "The Science and Joy of Baking", Content: "Baking transforms simple ingredients into delightful treats. From cookies to cakes, the process is both science and art, rooted in tradition.\n\nSharing baked goods with loved ones strengthens bonds. Experimenting with flavors like lavender or matcha creates memorable desserts.", Tags: []string{"Baking", "Cooking", "Desserts"}},
	{Title: "Conducting Science Experiments at Home", Content: "Science experiments fuel discovery by testing ideas against reality. Simple projects, like building circuits or observing chemical reactions, ignite curiosity.\n\nLabs or home setups teach precision and critical thinking. Sharing findings with peers drives innovation and collaboration.", Tags: []string{"Science", "Experiments", "Discovery"}},
	{Title: "The Beauty of Poetry in Expression", Content: "Poetry weaves emotions into concise, powerful verses. Crafting metaphors or free verse captures experiences in ways prose cannot.\n\nPerforming at open mics or publishing in journals connects poets with audiences. Reading classics inspires fresh, authentic work.", Tags: []string{"Poetry", "Writing", "Expression"}},
	{Title: "Cycling for Fitness and Freedom", Content: "Cycling offers freedom and fitness on two wheels. Whether commuting or exploring trails, pedaling builds endurance and reduces stress.\n\nGroup rides or charity events foster camaraderie. Maintaining your bike ensures safe, enjoyable adventures.", Tags: []string{"Cycling", "Fitness", "Adventure"}},
	{Title: "Unlocking Cultures Through Language Learning", Content: "Learning languages unlocks global communication and cultural understanding. Immersion through travel or apps like Duolingo accelerates fluency.\n\nPracticing with native speakers builds confidence. Language exchanges or media consumption deepen cultural appreciation.", Tags: []string{"Languages", "Culture", "Learning"}},
	{Title: "The Principles of Modern Architecture", Content: "Architecture blends form and function to create inspiring spaces. From skyscrapers to homes, designs reflect cultural and practical needs.\n\nStudying blueprints or visiting iconic buildings reveals their stories. Urban planning shapes sustainable, livable cities.", Tags: []string{"Architecture", "Design", "Urban Planning"}},
	{Title: "The Harmony of Singing and Music", Content: "Singing channels emotions through melody and harmony. Whether in choirs or solo, vocal practice builds confidence and technical skill.\n\nPerforming at events or recording covers shares your voice. Exploring genres from opera to pop expands versatility.", Tags: []string{"Singing", "Music", "Performance"}},
	{Title: "Strategic Thinking with Chess", Content: "Chess sharpens strategic thinking through calculated moves. Each game teaches anticipation and adaptability, rewarding patience and foresight.\n\nJoining clubs or online platforms like Chess.com hones skills. Studying grandmaster games inspires new tactics.", Tags: []string{"Chess", "Strategy", "Gaming"}},
	{Title: "Handmade Crafts and Creativity", Content: "Crafts turn raw materials into unique creations. From woodworking to sewing, hands-on projects spark creativity and satisfaction.\n\nSharing work at markets or online builds community. Tutorials on platforms like YouTube refine techniques.", Tags: []string{"Crafts", "Creativity", "DIY"}},
	{Title: "Philosophical Explorations of Life", Content: "Philosophy explores life’s big questions, from ethics to existence. Engaging with thinkers like Plato or modern debates sharpens critical thinking.\n\nDiscussing ideas in forums or reading essays fosters clarity. Applying philosophy to daily decisions adds depth to choices.", Tags: []string{"Philosophy", "Ethics", "Critical Thinking"}},
	{Title: "Swimming Techniques and Benefits", Content: "Swimming builds strength and endurance in a low-impact environment. Mastering strokes like freestyle or butterfly enhances technique and confidence.\n\nCompeting in meets or swimming in open waters offers thrill and challenge. Pools or beaches provide refreshing escapes.", Tags: []string{"Swimming", "Fitness", "Sports"}},
	{Title: "The Ethics of Journalism", Content: "Journalism uncovers truths through rigorous reporting. Interviewing sources and verifying facts ensure stories resonate with accuracy and impact.\n\nPublishing articles or breaking news informs communities. Ethical reporting builds trust and drives change.", Tags: []string{"Journalism", "Reporting", "News"}},
	{Title: "The Magic of Theater Performances", Content: "Theater brings stories to life through performance and staging. Actors embody characters, creating emotional connections with audiences.\n\nRehearsals refine timing and delivery. Attending or producing plays celebrates the art of storytelling.", Tags: []string{"Theater", "Performance", "Storytelling"}},
	{Title: "Building and Programming Robots", Content: "Robotics combines engineering and programming to create intelligent machines. Building bots for tasks or competitions sparks innovation.\n\nSensors and code bring designs to life. Joining robotics clubs or hackathons accelerates learning and creativity.", Tags: []string{"Robotics", "Engineering", "Programming"}},
	{Title: "The Art of Knitting Creations", Content: "Knitting creates cozy, handmade items with simple tools. Patterns for scarves or sweaters guide stitches into functional art.\n\nSharing projects at craft fairs or online builds community. Mastering techniques like cables adds complexity.", Tags: []string{"Knitting", "Crafts", "DIY"}},
	{Title: "Understanding Economic Markets", Content: "Economics analyzes how resources shape societies. Understanding markets or policies helps predict trends and inform decisions.\n\nStudying data or historical trends sharpens insights. Applying economic principles aids personal and societal growth.", Tags: []string{"Economics", "Markets", "Policy"}},
	{Title: "Capturing Landscapes in Painting", Content: "Painting landscapes captures the beauty of nature on canvas. Blending colors and techniques creates vivid, emotional scenes.\n\nExhibiting in galleries or online shares your vision. Studying masters like Monet inspires new approaches.", Tags: []string{"Painting", "Art", "Nature"}},
	{Title: "Producing Engaging Podcasts", Content: "Podcasts deliver stories and expertise to global listeners. Producing episodes on niche topics builds engaged audiences.\n\nEditing with tools like Audacity polishes content. Guest interviews add depth and spark dynamic conversations.", Tags: []string{"Podcasts", "Media", "Storytelling"}},
	{Title: "Adventurous Hiking Trails", Content: "Hiking offers adventure and connection to nature. Trails through forests or mountains challenge fitness and reward with stunning views.\n\nPacking essentials and respecting wildlife ensures safety. Sharing routes or photos inspires others to explore.", Tags: []string{"Hiking", "Nature", "Adventure"}},
	{Title: "Designing User-Friendly Interfaces", Content: "Design shapes user-friendly interfaces for apps and websites. Thoughtful layouts and colors enhance usability and engagement.\n\nPrototyping tools like Figma streamline testing. Launching designs that solve real problems delights users.", Tags: []string{"Design", "UI/UX", "Technology"}},
	{Title: "The Grace of Ballet Dance", Content: "Ballet combines grace and discipline in every movement. Pirouettes and pliés demand precision, building strength and poise.\n\nPerforming in recitals or studying classics inspires growth. Classes foster a supportive dance community.", Tags: []string{"Ballet", "Dance", "Performance"}},
	{Title: "Discovering the Cosmos in Astronomy", Content: "Astronomy reveals the universe’s mysteries, from black holes to exoplanets. Observing with telescopes or apps fuels wonder and discovery.\n\nFollowing NASA missions or joining star parties deepens knowledge. The cosmos invites endless exploration.", Tags: []string{"Astronomy", "Space", "Science"}},
	{Title: "Sculpting Expressive Forms", Content: "Sculpting transforms clay or stone into expressive forms. Molding materials by hand connects artists to their creations.\n\nExhibiting sculptures in galleries shares stories. Learning techniques like carving refines artistic vision.", Tags: []string{"Sculpting", "Art", "Creativity"}},
	{Title: "Honing Skills Through Debate", Content: "Debate hones critical thinking and persuasive skills. Crafting arguments backed by evidence prepares you for any challenge.\n\nParticipating in tournaments or forums builds confidence. Studying rhetoric sharpens communication.", Tags: []string{"Debate", "Argumentation", "Communication"}},
	{Title: "The Intricate Art of Origami", Content: "Origami turns paper into intricate art through precise folds. Creating cranes or flowers teaches patience and focus.\n\nSharing models at workshops or online inspires others. Mastering complex designs elevates the craft.", Tags: []string{"Origami", "Crafts", "Art"}},
	{Title: "Navigating the Seas with Sailing", Content: "Sailing navigates open waters with skill and intuition. Harnessing wind and waves offers freedom and adventure.\n\nLearning knots or charting courses builds expertise. Coastal or ocean voyages create unforgettable memories.", Tags: []string{"Sailing", "Adventure", "Navigation"}},
	{Title: "Bringing Stories to Life with Illustration", Content: "Illustration brings stories to life with vivid drawings. Sketching characters or scenes adds depth to books or media.\n\nUsing digital tools or traditional pencils, illustrators capture imagination. Publishing in children's books or graphic novels reaches wide audiences.", Tags: []string{"Illustration", "Art", "Storytelling"}},
	{Title: "Achieving Clarity Through Meditation", Content: "Meditation centers minds and reduces daily stresses. Simple breathing exercises anchor presence and promote inner peace.\n\nRegular sessions, even brief ones, yield profound insights and emotional balance. Integrating mindfulness into routines enhances overall well-being.", Tags: []string{"Meditation", "Mindfulness", "Wellness"}},
	{Title: "The Evolution of Digital Media", Content: "Digital media continues to transform how we consume and create content. From social platforms to streaming services, accessibility has never been greater.\n\nNavigating this landscape requires media literacy to discern quality from noise. Encouraging diverse voices ensures a richer, more inclusive digital ecosystem.", Tags: []string{"Digital Media", "Content Creation", "Literacy"}},
	{Title: "Sustainable Practices in Everyday Life", Content: "Adopting sustainable practices can significantly reduce our environmental footprint. Simple swaps like reusable bags or energy-efficient appliances add up over time.\n\nCommunity initiatives amplify individual efforts, fostering collective action. Education on climate impacts empowers informed, proactive choices.", Tags: []string{"Sustainability", "Environment", "Lifestyle"}},
	{Title: "The Science Behind Human Behavior", Content: "Understanding human behavior through psychology reveals patterns in decision-making and emotions. Insights from studies help improve relationships and self-awareness.\n\nApplying these principles in therapy or coaching leads to meaningful change. Ongoing research keeps the field dynamic and relevant to modern challenges.", Tags: []string{"Psychology", "Behavior", "Self-Awareness"}},
	{Title: "Innovations in Renewable Energy", Content: "Renewable energy sources like solar and wind are key to a greener future. Technological advances make them more efficient and cost-effective.\n\nPolicy support and investment accelerate adoption worldwide. Transitioning from fossil fuels benefits both the economy and the planet.", Tags: []string{"Renewable Energy", "Innovation", "Environment"}},
	{Title: "Cultivating Mindfulness in a Busy World", Content: "Mindfulness practices counteract the rush of modern life, promoting focus and gratitude. Techniques like guided meditation offer quick resets.\n\nIncorporating them into daily routines builds resilience against stress. Long-term benefits include improved mental health and productivity.", Tags: []string{"Mindfulness", "Stress Management", "Wellness"}},
	{Title: "The Future of Artificial Intelligence", Content: "Artificial intelligence is poised to revolutionize sectors from healthcare to transportation. Ethical development ensures it augments rather than replaces human capabilities.\n\nCollaboration between tech experts and policymakers will shape its trajectory. Embracing AI thoughtfully unlocks unprecedented opportunities.", Tags: []string{"AI", "Technology", "Ethics"}},
	{Title: "Exploring Urban Gardening Techniques", Content: "Urban gardening maximizes limited spaces with vertical planters and hydroponics. It provides fresh produce and a therapeutic hobby.\n\nCommunity gardens strengthen neighborhood ties and promote food security. Starting small encourages sustainable habits that last.", Tags: []string{"Urban Gardening", "Sustainability", "Community"}},
}

var comments = []string{
	"Thanks for sharing, really interesting perspective!", "Wow, I learned something new today. Great stuff!", "This was super helpful, appreciate the details.", "Love the enthusiasm in this! Keep it up.", "Such a great read, got me thinking.",
	"Amazing content, definitely sharing this!", "Thanks for the tips, I’ll try them out.", "This is so inspiring, love your approach!", "Really enjoyed this, more like this please!", "Great job breaking this down, super clear.", "I’m definitely bookmarking this for later.", "Fantastic insights, thanks for posting!",
	"This was a fun read, well done!", "Such a cool topic, learned a lot here.", "Thanks for the advice, super practical!", "Really engaging, can’t wait for more!", "This made my day, awesome content!", "Super informative, thanks for the effort!", "Great points, I totally agree!",
	"This was so well explained, thank you!", "Love how you presented this, very clear!", "Thanks for the ideas, I’m inspired!", "Really thoughtful post, enjoyed it!", "Great content, keep up the good work!", "This was eye-opening, thanks for sharing!", "Super useful info, I’ll use this!",
	"Love the creativity here, nicely done!", "Thanks for the insights, really valuable!", "This was a great take, learned something!", "Appreciate the effort, super engaging!", "Such a cool perspective, thanks for this!", "Well written, really enjoyed reading it!", "Thanks for the tips, very actionable!",
	"Great stuff, I’m sharing this with friends!", "This was so interesting, great job!", "Love the energy in this, keep it coming!", "Really helpful advice, thank you!", "Fantastic read, super thought-provoking!", "Thanks for breaking this down, awesome!",
	"This was such a fun post to read!", "Great ideas here, I’m taking notes!", "Super insightful, thanks for sharing!", "Love how clear this is, great work!", "This was really motivating, thank you!", "Such a great post, learned a ton!", "Thanks for the info, really useful!", "Really enjoyed this, well done!", "Great content, super engaging!",
	"This was so helpful, thanks a lot!", "Love the ideas here, very inspiring!",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(500)
	for i, user := range users {
		fmt.Println("Creating user: ", i+1)
		err := store.Users.Create(ctx, user)
		if err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}

	posts := generatePosts(10000, users)
	for i, post := range posts {
		fmt.Println("Creating post: ", i+1)
		err := store.Posts.Create(ctx, post)
		if err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(100000, users, posts)
	for i, comment := range comments {
		fmt.Println("Creating comment: ", i+1)
		err := store.Comments.Create(ctx, comment)
		if err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}

	// generate followers
	followers := 100000
	for i := range followers {
		fmt.Println("Creating follower: ", i+1)
		followerID := rand.IntN(len(users))
		userID := rand.IntN(len(users))
		err := store.Followers.Follow(ctx, int64(followerID), int64(userID))
		if err != nil {
			log.Printf("Error creating follower (%d): %s\n", i+1, err)
		}
	}

}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := range num {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    emails[i%len(emails)] + fmt.Sprintf("_%d@example.com", i),
			Password: "123",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := range num {
		posts[i] = &store.Post{
			Title:   postsMock[i%len(postsMock)].Title,
			UserID:  users[rand.IntN(len(users))].ID,
			Content: postsMock[i%len(postsMock)].Content,
			Tags:    postsMock[i%len(postsMock)].Tags,
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)

	for i := range num {
		cms[i] = &store.Comment{
			PostID:  posts[rand.IntN(len(posts))].ID,
			UserID:  users[rand.IntN(len(users))].ID,
			Content: comments[i%len(comments)],
		}
	}

	return cms
}
