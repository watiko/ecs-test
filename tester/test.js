const got = require('got');

const PORT = process.env.PORT || 8080;

async function wait(sec) {
  while (true) {
    const res = await got(`http://localhost:${PORT}/wait/${sec}`);
    console.log(res.body);
  }
}

async function main() {
  const waits = [5, 15, 30, 60, 60 * 2, 60 * 3];
  
  await Promise.all(waits.map(wait));
}

main()
  .catch(e => console.error(e))
