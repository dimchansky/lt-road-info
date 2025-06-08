# Security Policy

## Security Considerations

This project downloads and processes road information from public Lithuanian government APIs to generate GPX files for navigation purposes. While the application has a minimal security attack surface, we take security seriously.

## Supported Versions

We provide security updates for the following versions:

| Version | Supported          |
| ------- | ------------------ |
| Latest  | :white_check_mark: |
| < 1.0   | :white_check_mark: |

## Security Model

### Data Sources
- **Public APIs Only**: This application only accesses publicly available Lithuanian government APIs
- **Read-Only Operations**: No data is written to external systems
- **HTTPS Only**: All API communications use encrypted HTTPS connections

### Generated Files
- **Local File System**: GPX files are only written to the local file system
- **No Network Exposure**: Generated files are not automatically uploaded or shared
- **Standard File Permissions**: Files are created with standard read/write permissions (0644)

### Dependencies
- **Minimal Dependencies**: Uses a small set of well-maintained Go libraries
- **Coordinate Transformation**: Uses specialized libraries for accurate geographic calculations
- **HTTP Client**: Uses Go's standard HTTP client with default security settings

## Known Security Considerations

### 1. Network-Based Attacks
- **Man-in-the-Middle**: HTTPS provides protection, but certificate validation relies on system trust stores
- **DNS Attacks**: Domain resolution relies on system DNS configuration

### 2. File System
- **Output Directory**: Users must ensure appropriate permissions on output directories
- **File Overwrites**: The application will overwrite existing GPX files with the same name

### 3. Data Integrity
- **API Response Validation**: The application validates JSON structure but relies on API providers for data integrity
- **Coordinate Accuracy**: Geographic accuracy is validated through comprehensive testing

## Reporting Security Vulnerabilities

### For Security Issues
If you discover a security vulnerability, please report it privately:

**Email**: [Create a GitHub issue and mention it's security-related]
**GitHub**: Use private vulnerability reporting if available

### What to Include
Please include the following information:
- Description of the vulnerability
- Steps to reproduce the issue
- Potential impact
- Suggested fix (if available)

### Response Timeline
- **Initial Response**: Within 48 hours
- **Assessment**: Within 7 days
- **Fix**: Depends on severity and complexity

## Security Best Practices for Users

### When Running the Application
1. **Verify Downloads**: Only download releases from the official GitHub repository
2. **Check Checksums**: Verify binary integrity if checksums are provided
3. **Network Security**: Ensure you're on a trusted network when downloading road data
4. **File Permissions**: Review permissions on generated GPX files

### When Contributing
1. **Code Review**: All code changes are reviewed before merging
2. **Dependency Updates**: Keep dependencies updated to their latest secure versions
3. **Static Analysis**: Use `go vet` and security linters when available

### For System Administrators
1. **Firewall Rules**: Allow HTTPS connections to `eismoinfo.lt` and `gis.ktvis.lt` domains
2. **Output Directory**: Ensure appropriate directory permissions for GPX file output
3. **Log Monitoring**: Monitor for unusual network activity or file system access

## Security Features

### Built-in Protections
- **Input Validation**: JSON responses are validated and parsed safely
- **Error Handling**: Failures are handled gracefully without exposing sensitive information
- **Memory Safety**: Go's memory management prevents common vulnerabilities
- **No Eval**: No dynamic code execution or user input evaluation

### Testing Security
- **Static Analysis**: Code is checked with `go vet`
- **Dependency Scanning**: Dependencies are monitored for known vulnerabilities
- **Geographic Validation**: Prevents coordinate manipulation that could mislead users

## Threat Model

### In Scope
- Vulnerabilities in application code
- Insecure dependency usage
- File system security issues
- Network communication security

### Out of Scope
- Vulnerabilities in underlying operating system
- DNS or network infrastructure attacks
- Physical access to the machine
- Social engineering attacks

## Incident Response

In case of a confirmed security vulnerability:

1. **Assessment**: Evaluate the severity and potential impact
2. **Fix Development**: Develop and test a security patch
3. **Coordinated Disclosure**: Work with the reporter to coordinate public disclosure
4. **Release**: Issue a security update with clear documentation
5. **Notification**: Notify users through GitHub releases and README updates

## Contact

For security-related questions or concerns:
- **GitHub Issues**: For non-sensitive security discussions
- **Private Reporting**: Use GitHub's private vulnerability reporting feature

## Updates

This security policy may be updated as the project evolves. Check back regularly for the latest information.

---

**Note**: This project processes public road information and generates navigation files. While we implement security best practices, users should always follow official road signs and regulations, as the primary safety concern is accurate navigation rather than information security.