# Exploratory Mongo db layout

db.words.deleteMany({})
db.vocab.deleteMany({})
db.users.deleteMany({})

db.vocab.find({})
db.words.find({})
db.users.find({})
db.createCollection("words")
db.createCollection("users")
db.createCollection("vocab")

db.words.createIndex( { key:1, language:1}, { unique: true } )
db.words.createIndex( { id:1 }, { } )
db.words.insert({id:1,language:"ja", src:"jisho", key:"猫", alts:["ねこ"], data: {meaning:"cat"}})
db.words.insert({id:2,language:"ja", src:"jisho", key:"犬", alts:["いぬ"], data: {meaning:"dog"}})

db.users.createIndex( { email:1}, { unique: true } )
db.users.insert({name:"a", languages:["ja", "ru"], lastLogin:"", email:"e1"})
db.users.insert({name:"b", languages:["ja", "en"], lastLogin:"", email:"e2"})

db.vocab.createIndex( { word_id:1, user:1 }, { unique: true } )
db.vocab.createIndex( { user:1, user:1 }, { } )
db.vocab.insert({word_id:1, user:"e1", language:"ja", alts:["犬", "いぬ"], date:"", seen:1, confidence:6})
db.vocab.insert({word_id:2, user:"e1", language:"ja", alts:["猫", "ねこ"], date:"", seen:1, confidence:6})

db.users.aggregate([
  { $match : { email : 'e1' } },
  { $lookup : {
      from : 'vocab',
      localField : 'email',
      foreignField : 'user',
      as : 'vocab'
  } }
]).pretty()
    
    