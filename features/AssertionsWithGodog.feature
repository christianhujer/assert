Feature: Using Assertions with Godog
  As a user of Godog
  I want to be able to use the assert library
  So that I can use the same assertions as in testing

  Scenario:
    Given I have 42 hotdogs
    When I eat 19 hotdogs
    Then I MUST have 23 hotdogs