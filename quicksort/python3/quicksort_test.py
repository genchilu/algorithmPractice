import unittest
import quicksort

class QuicksortTestCase(unittest.TestCase):

    def test_sort(self):
        test_cases = [
            {'input': [], 'expect': []},
            {'input': None, 'expect': None},
            {'input': [1], 'expect': [1]},
            {'input': [1,2], 'expect': [1,2]},
            {'input': [2,1], 'expect': [1,2]},
            {'input': [1,1], 'expect': [1,1]},
            {'input': [1,2,3,4,5], 'expect': [1,2,3,4,5]},
            {'input': [5,4,3,2,1], 'expect': [1,2,3,4,5]},
            {'input': [7,5,3,3,9,6,7], 'expect': [3,3,5,6,7,7,9]},
            {'input': [23, 42, 32, 64, 12, 4], 'expect': [4, 12, 23, 32, 42, 64]},
            {'input': [2,1,2,2,2,2,2], 'expect': [1,2,2,2,2,2,2]},
            {'input': [10,12,9,13,8,14,7], 'expect': [7,8,9,10,12,13,14]},
        ]

        for test_case in test_cases:
            quicksort.sort(test_case['input']);
            self.assertEqual(test_case['expect'], test_case['input']);

if __name__ == '__main__':unittest.main()