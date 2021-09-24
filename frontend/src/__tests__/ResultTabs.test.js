import React from "react";
import ResultTabs from "../components/ResultTabs";
import { shallow } from 'enzyme';

it('Tab is changed when tab is clicked', () => {
  const wrapper = shallow(
       <ResultTabs />
  );

  // prints tabs component
  console.log(wrapper.find('#tabs').debug());
  expect(wrapper).toBeDefined()

});