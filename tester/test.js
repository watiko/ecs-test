const got = require('got');
const { Agent } = require('http');

const PORT = process.env.PORT || 8080;

const sleep = ms => new Promise(resolve => {
  setTimeout(resolve, ms);
})

const jitteredSleep = maxSec => sleep(Math.random() * maxSec * 1000);

async function reqWait(sleepSec, waitSec) {
  const options = {
    agent: {
      http: new Agent({ keepAlive: true }),
    },
  };

  while (true) {
    await jitteredSleep(10);
    try {
      const res = await got(`http://localhost:${PORT}/wait/${waitSec}`, options);
      console.log(res.body);
      await sleep(sleepSec * 1000);
    } catch (err) {
      console.error('failed to request', err);
    }
  }
}

async function main() {
  let params = [[20, 30]];

  for (let i = 0; i < 100; i++) {
    params = [[1, 1], [2, 0.45], [3, 0.45], [4, 0.45], ...params]
  }
  
  await Promise.all(params.map(([sleepSec, waitSec]) => reqWait(sleepSec, waitSec)));
}

main()
  .catch(e => console.error(e))
