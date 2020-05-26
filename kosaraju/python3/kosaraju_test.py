import unittest
import kosaraju

class KosarajuTestCase(unittest.TestCase):

    def test_do_kosaraju_scc(self):
        test_cases = [
            {
                'inputEdges': [
                    [6,7],
                    [7,0],
                    [7,4],
                    [7,5],
                    [5,6],
                    [5,4],
                    [4,3],
                    [3,2],
                    [3,4],
                    [0,1],
                    [0,2],
                    [0,3],
                    [1,0],
                    [1,2],
                ], 
                'expectComponments': [
                    [5,6,7],
                    [3,4],
                    [2],
                    [0,1]
                ]
            },
            {
                'inputEdges': [
                    [1,0],
                    [2,1],
                    [0,2],
                    [0,3],
                    [3,4],
                ], 
                'expectComponments': [
                    [0,1,2],
                    [3],
                    [4],
                ]
            },
            {
                'inputEdges': [
                    [1,3],
                    [2,1],
                    [3,2],
                    [3,4],
                    [4,5],
                    [5,6],
                    [6,4],
                    [7,8],
                    [8,10],
                    [9,7],
                    [9,6],
                    [10,9],
                    [10,11],
                ], 
                'expectComponments': [
                    [1,2,3],
                    [4,5,6],
                    [7,8,9,10],
                    [11]
                ]
            },
        ]

        for test_case in test_cases:
            actualComponments = kosaraju.do_kosaraju_scc(test_case['inputEdges']);
            
            self.assertEqual(len(test_case['expectComponments']), len(actualComponments))
            for expectComponment in test_case['expectComponments']:
                lowest_value = expectComponment[0]
                componment = actualComponments[lowest_value]

                self.assertEqual(len(expectComponment), len(componment))
                
                i = 0
                while i < len(expectComponment):
                    self.assertEqual(expectComponment[i], componment[i])
                    i = i + 1

if __name__ == '__main__':unittest.main()