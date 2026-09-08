import os
from pathlib import Path
import subprocess
import tempfile
import unittest

SCRIPT = Path(__file__).resolve().parents[1] / "generate-java.sh"
GRAMMAR = SCRIPT.parent / "MojoParser.g4"


class GenerateJavaTest(unittest.TestCase):
    def test_success_and_failure_preserve_grammar(self):
        original = GRAMMAR.read_bytes()
        with tempfile.TemporaryDirectory(prefix="mojo java ") as directory:
            root = Path(directory)
            bin_dir = root / "bin"
            bin_dir.mkdir()
            runner = bin_dir / "antlr4"
            runner.write_text("""#!/usr/bin/env python3
import pathlib, sys
args = sys.argv[1:]
assert '-Dlanguage=Java' in args
assert '-no-listener' in args
text = pathlib.Path('MojoParser.g4').read_text()
assert 'p.GetTokenStream()' not in text
assert '_input.index() > 0 && _input.get(_input.index()-1).getType() != WS' in text
out = pathlib.Path(args[args.index('-o') + 1])
out.mkdir(parents=True)
for name in ['MojoLexer', 'MojoParser', 'MojoParserBaseVisitor', 'MojoParserVisitor']:
    (out / (name + '.java')).write_text('// generated for script test')
""")
            runner.chmod(0o755)
            env = os.environ.copy()
            env.pop('ANTLR_JAR', None)
            env['PATH'] = str(bin_dir) + os.pathsep + env['PATH']
            output = root / 'output with spaces'
            result = subprocess.run(['bash', str(SCRIPT), str(output)], cwd=root, env=env, capture_output=True, text=True)
            self.assertEqual(0, result.returncode, result.stderr)
            self.assertEqual(4, len(list(output.glob('*.java'))))
            self.assertEqual(original, GRAMMAR.read_bytes())
            runner.write_text('#!/bin/sh\nexit 42\n')
            result = subprocess.run(['bash', str(SCRIPT), str(output)], cwd=root, env=env, capture_output=True, text=True)
            self.assertEqual(42, result.returncode)
            self.assertEqual(original, GRAMMAR.read_bytes())


if __name__ == '__main__':
    unittest.main()
