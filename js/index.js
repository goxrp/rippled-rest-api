const express = require('express') // npm install express --save
const app = express()
const port = 3000
const axios = require('axios').default; // npm install axios --save

app.use(express.json());

app.get('/', (req, res) => {
    res.send('I am alive!')
})

app.post('/api/:method', async function (req, res) {
    apiReqBody = {
        'method': req.params['method'],
        'params': [req.body]
    }
    rpcUrl = 'https://s1.ripple.com:51234/'
    try {
        const apiResBody = await axios.post(rpcUrl, apiReqBody)
        console.log(JSON.stringify(apiResBody.data))
        res.send(apiResBody.data.result)
    } catch (error) {
        console.log("API_ERROR")
        console.log(Object.keys(error), error.message);
    }
})

app.listen(port, () => {
    console.log(`Example app listening at http://localhost:${port}`)
})


/*

curl -XPOST https://s1.ripple.com:51234/ -H 'Content-Type: application/json' -d '{"method":"account_info","params":[{"account":"rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn","strict":true,"ledger_index":"current","queue":true}]}'



curl -XPOST https://s1.ripple.com:51234/ -H 'Content-Type: application/json' -d '{"method":"account_info","params":[{"account":"rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn","strict":true,"ledger_index":"current","queue":true}]}'


{"method":"account_info","params":[{"account":"rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn","strict":true,"ledger_index":"current","queue":true}]}


curl -XPOST http://localhost:3000/api/account_info -H 'Content-Type: application/json' -d '{"account": "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn", "strict": true, "ledger_index": "current", "queue": true}'


{"method": "account_info","params": [       {
            "account": "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
            "strict": true,
            "ledger_index": "current",
            "queue": true
        }
    ]
}
*/