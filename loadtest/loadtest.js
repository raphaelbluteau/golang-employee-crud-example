import http from 'k6/http';
import { sleep } from 'k6';

export let options = {
    stages: [
        { duration: '2m', target: 1 }, // simulate ramp-up of traffic from 1 to 200 users over 2 minutes
        { duration: '1m', target: 200 }, // stay at 200 users for 1 minutes
        { duration: '1m', target: 0 }, // ramp-down to 0 users
    ],
    thresholds: {
        'http_req_duration': ['p(99)<1500'], // 99% of requests must complete below 1.5s
    },
};

export function testSuite() {
    const BASE_URL = 'http://app:1323';

    // Create Employee
    let url = `${BASE_URL}/employees`;
    let payload = JSON.stringify({
        first_name: 'John',
        last_name: 'Doe',
        email: 'johndoe@example.com',
    });
    let params = { headers: { 'Content-Type': 'application/json' } };
    let res = http.post(url, payload, params);
    let id = JSON.parse(res.body).id;

    // Get Employee
    url = `${BASE_URL}/employees/${id}`;
    res = http.get(url);

    // Update Employee
    url = `${BASE_URL}/employees/${id}`;
    payload = JSON.stringify({
        first_name: 'Jane',
        last_name: 'Doe',
        email: 'janedoe@example.com',
    });
    res = http.put(url, payload, params);

    // Delete Employee
    url = `${BASE_URL}/employees/${id}`;
    res = http.del(url);
}

// Ramp-up and ramp-down can generate a lot of HTTP requests,
// and k6 has a default maximum number of open file descriptors (ulimit).
// This can cause k6 to crash with the "too many open files" error.
// To prevent this, we increase the maximum number of file descriptors to 300000.
// You can adjust this value as needed for your specific test.
export function setup() {
    const maxFileDescriptors = 300000;
    return {
        'ulimit -n': maxFileDescriptors,
    };
}

export default testSuite;
