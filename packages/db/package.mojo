/// the db package
package mojo.db {
    version: '0.1.0'
    authors: [{
        author: 'frankee'
        email:  'frankee.zhou@gmail.com'
        organization: 'mojolang.org'
    }]
    
    dependencies: {
        'mojo.core': {repository: 'github.com/mojo-lang/mojo/packages/core', version: '^0.1'}
        'mojo.document': {repository: 'github.com/mojo-lang/mojo/packages/document', version: '^0.1'}
        'mojo.lang': {repository: 'github.com/mojo-lang/mojo/packages/lang', version: '^0.1'}
    }
    repository: 'github.com/mojo-lang/mojo/packages/db'
}