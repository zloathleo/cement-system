import MockAdapter from 'axios-mock-adapter'
import Mock from 'mockjs'
import myaxios from './myaxios';

let mock = new MockAdapter(myaxios, {
    delayResponse: 300
});

mock.onGet('/dashboard-commondata').reply(200, Mock.mock(
    {
        d1: 100,
        d2: 0,
        d3: 0,
        d4: 1,
    }
));

mock.onGet('/dashboard-data').reply(200, Mock.mock({
    "def":
        {
            label: 'def',
            beginTime: 0,
            endTime: 0,
            status: 1,
            updateTime: 0,
        }
}));