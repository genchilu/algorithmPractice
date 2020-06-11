import unittest
import karger_mincut

class KargerMinCutTestCases(unittest.TestCase):
    def test_do_min_cut(self):
        test_cases = [
            {
                'input': None,
                'expect': 0 
            },
            {
                'input': [],
                'expect': 0 
            },
            {
                'input':  [
                    [1,2],
                    [2,3],
                    [0,4],
                    [1,5],
                    [2,6],
                    [3,7],
                    [4,5],
                    [5,6],
                    [6,7],
                    [0,5],
                    [1,4],
                    [2,7],
                    [3,6],
                ],
                'expect': 2
            },
            {
                'input':  [
                    [0,1],
                    [0,2],
                    [0,3],
                    [0,4],
                    [1,2],
                    [1,3],
                    [1,4],
                    [2,3],
                    [2,4],
                    [3,4],
                    [5,6],
                    [5,7],
                    [5,8],
                    [5,9],
                    [6,7],
                    [6,8],
                    [6,9],
                    [7,8],
                    [7,9],
                    [8,9],
                    [0,5],
                    [2,6],
                    [4,7]
                ],
                'expect': 3
            }]

        for test_case in test_cases:
            actual_min_cut = karger_mincut.do_min_cut(test_case['input'])
            self.assertEqual(test_case['expect'], actual_min_cut)

if __name__ == '__main__':unittest.main()
