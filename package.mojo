// Standard library packages share the source and output trees of this repository.

/// the core package
package mojo.core {
    version: '0.1.0'
    authors: [{
        author: 'Frankee'
        email: 'frankee.zhou@gmail.com'
        organization: 'mojolang.org'
    }]

    license: 'Apache'
    repository: 'github.com/mojo-lang/mojo'
}

///
package mojo.document {
    version: '0.1.0'
    license: ''
    authors: [{
        author: 'Frankee'
        email: 'frankee.zhou@gmail.com'
        organization: 'mojolang.org'
    }]
    dependencies: {
        'mojo.core': {repository: 'github.com/mojo-lang/mojo', version: '^0.1'}
    }

    repository: 'https://github.com/mojo-lang/mojo'
}

///
package mojo.lang {
    version: '0.1.0'
    license: 'Apache'
    authors: [{
        author: 'Frankee'
        email: 'frankee.zhou@gmail.com'
        organization: 'mojolang.org'
    }]
    dependencies: {
        'mojo.core': {repository: 'github.com/mojo-lang/mojo', version: '^0.1'}
        'mojo.document': {repository: 'github.com/mojo-lang/mojo', version: '^0.1'}
    }

    repository: 'https://github.com/mojo-lang/mojo'
}

/// the db package
package mojo.db {
    version: '0.1.0'
    authors: [{
        author: 'frankee'
        email:  'frankee.zhou@gmail.com'
        organization: 'mojolang.org'
    }]
    
    dependencies: {
        'mojo.core': {repository: 'github.com/mojo-lang/mojo', version: '^0.1'}
        'mojo.document': {repository: 'github.com/mojo-lang/mojo', version: '^0.1'}
        'mojo.lang': {repository: 'github.com/mojo-lang/mojo', version: '^0.1'}
    }
    repository: 'github.com/mojo-lang/mojo'
}

///
package mojo.geom {
    version: '0.1.0'
    license: ''
    authors: [{
        author: 'Frankee'
        email: 'frankee.zhou@gmail.com'
        organization: 'mojolang.org'
    }]
    dependencies: {
        'mojo.core': {repository: 'github.com/mojo-lang/mojo', version: '^0.1'}
    }

    repository: 'https://github.com/mojo-lang/mojo'
}

/// the http package
package mojo.http {
    version: '0.1.0'
    authors: [{
        author: 'frankee'
        email:  'frankee.zhou@gmail.com'
        organization: 'mojolang.org'
    }]
    
    dependencies: {
        'mojo.core': {repository: 'github.com/mojo-lang/mojo', version: '^0.1'}
    }
    repository: 'github.com/mojo-lang/mojo'
}

///
package mojo.openapi {
    version: '0.1.0'
    license: 'Apache'
    authors: [{
        author: 'Frankee'
        email: 'frankee.zhou@gmail.com'
        organization: 'mojolang.org'
    }]
    dependencies: {
        'mojo.core': {repository: 'github.com/mojo-lang/mojo', version: '^0.1'}
        'mojo.document': {repository: 'github.com/mojo-lang/mojo', version: '^0.1'}
    }

    repository: 'https://github.com/mojo-lang/mojo'
}

/// the rpc package
package mojo.rpc {
    version: '0.1.0'
    authors: [{
        author: 'frankee'
        email:  'frankee.zhou@gmail.com'
        organization: 'mojo-lang.org'
    }]
    
    dependencies: {
        'mojo.core': {repository: 'github.com/mojo-lang/mojo', version: '^0.1'}
    }
    
    repository: 'https://github.com/mojo-lang/mojo'
}

/// the protobuf package
package mojo.protobuf {
    version: '0.1.0'
    license: 'Apache'
    authors: [{
        author: 'Frankee'
        email: 'frankee.zhou@gmail.com'
        organization: 'mojolang.org'
    }]
    
    dependencies: {
        'mojo.core': {repository: 'github.com/mojo-lang/mojo', version: '^0.1'}
    }
    repository: 'https://github.com/mojo-lang/mojo'
}

///
package mojo.yaml {
    version: '0.1.0'
    license: 'Apache'
    authors: [{
        author: 'Frankee'
        email: 'frankee.zhou@gmail.com'
        organization: 'mojolang.org'
    }]
    dependencies: {
        'mojo.core': {repository: 'github.com/mojo-lang/mojo', version: '^0.1'}
    }

    repository: 'https://github.com/mojo-lang/mojo'
}

/// the net package
package mojo.net {
    version: '0.1.0'
    authors: [{
        author: ''
        email:  ''
        organization: ''
    }]
    
    dependencies: {
        'mojo.core': {repository: 'github.com/mojo-lang/mojo', version: '0.0.0-20211123010202-03f9f6e22fd2'}
        'mojo.geom': {repository: 'github.com/mojo-lang/mojo', version: '0.0.0-20211121061729-3cc68d7658b2'}
    }
    repository: 'github.com/mojo-lang/mojo'
}

